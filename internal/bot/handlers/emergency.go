package handlers

import (
	"t0t0dcyberbot/internal/bot/keyboards"
	"t0t0dcyberbot/internal/session"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type EmergencyHandler struct {
	Msg       MessagingService
	Emergency EmergencyClient
}

func NewEmergencyHandler(msg MessagingService, emergency EmergencyClient) *EmergencyHandler {
	return &EmergencyHandler{
		Msg:       msg,
		Emergency: emergency,
	}
}

func (h *EmergencyHandler) HandleEmergency(update tgbotapi.Update, sess *session.Session) tgbotapi.Chattable {
	sess.Step = session.StepContact

	text := h.Msg.FormatEmergencyIntro()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown

	if formURL := h.Emergency.FormURL(); formURL != "" {
		if kb := keyboards.EmergencyInline(formURL); kb != nil {
			msg.ReplyMarkup = *kb
		}
	}

	return msg
}

func (h *EmergencyHandler) HandleOpenForm(update tgbotapi.Update, sess *session.Session) tgbotapi.Chattable {
	sess.Step = session.StepContact

	text := h.Msg.FormatEmergencyWithForm()
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown

	return msg
}
