// internal/bot/handlers/start.go
package handlers

import (
	"t0t0dcyberbot/internal/bot/keyboards"
	"t0t0dcyberbot/internal/session"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StartHandler struct {
	Msg MessagingService
}

func (h *StartHandler) HandleStart(update tgbotapi.Update, sess *session.Session) tgbotapi.Chattable {
	sess.Reset()

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, h.Msg.FormatWelcome())
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.ReplyMarkup = keyboards.Main()
	return msg
}

func (h *StartHandler) HandleAbout(update tgbotapi.Update, sess *session.Session) tgbotapi.Chattable {
	sess.Step = session.StepIdle

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, h.Msg.FormatAbout())
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.ReplyMarkup = keyboards.Main()
	return msg
}
