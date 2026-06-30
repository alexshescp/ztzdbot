package handlers

import (
	"strings"

	"t0t0dcyberbot/internal/bot/keyboards"
	"t0t0dcyberbot/internal/session"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IncidentsHandler struct {
	Msg       MessagingService
	Emergency EmergencyClient
}

func (h *IncidentsHandler) HandleStartIncident(update tgbotapi.Update, sess *session.Session) tgbotapi.Chattable {
	sess.Reset()
	sess.Step = session.StepIncidentType

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, h.Msg.FormatIncidentStart())
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.ReplyMarkup = *keyboards.IncidentTypes()

	return msg
}

func (h *IncidentsHandler) HandleSelectType(update tgbotapi.Update, sess *session.Session, incidentType string) tgbotapi.Chattable {
	sess.Step = session.StepIncidentDescription
	sess.IncidentType = incidentType

	chatID := update.CallbackQuery.Message.Chat.ID
	text := h.Msg.FormatAskIncidentDescription(incidentType)

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	return msg
}

func (h *IncidentsHandler) HandleStep(update tgbotapi.Update, sess *session.Session) tgbotapi.Chattable {
	chatID := update.Message.Chat.ID
	text := strings.TrimSpace(update.Message.Text)

	switch sess.Step {

	case session.StepIncidentDescription:
		sess.IncidentText = text
		sess.Step = session.StepContact
		msg := tgbotapi.NewMessage(chatID,
			h.Msg.FormatAskContact(sess.IncidentType, sess.IncidentText))
		msg.ParseMode = tgbotapi.ModeMarkdown
		return msg

	case session.StepContact:
		sess.ContactText = text
		_ = h.Emergency.SubmitRequest(sess.ContactText, sess.IncidentText)

		resp := h.Msg.FormatIncidentReceived(sess.IncidentType, sess.IncidentText, sess.ContactText)
		sess.Reset()

		msg := tgbotapi.NewMessage(chatID, resp)
		msg.ParseMode = tgbotapi.ModeMarkdown
		msg.ReplyMarkup = keyboards.Main()
		return msg

	default:
		sess.Reset()
		msg := tgbotapi.NewMessage(chatID,
			"Начнём заново. Используйте кнопки.\n\n🧨 Сообщить об инциденте\n🆘 Экстренная поддержка")
		msg.ParseMode = tgbotapi.ModeMarkdown
		msg.ReplyMarkup = keyboards.Main()
		return msg
	}
}
