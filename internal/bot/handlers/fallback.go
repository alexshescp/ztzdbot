package handlers

import (
	"t0t0dcyberbot/internal/session"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type FallbackHandler struct {
	Msg MessagingService
}

func NewFallbackHandler(msg MessagingService) *FallbackHandler {
	return &FallbackHandler{Msg: msg}
}

func (h *FallbackHandler) HandleUnknown(update tgbotapi.Update, sess *session.Session) tgbotapi.Chattable {
	text := "Я пока не понял запрос. Пожалуйста, воспользуйтесь кнопками на клавиатуре: «🧨 Сообщить об инциденте» или «🆘 Экстренная поддержка»."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = tgbotapi.ModeMarkdown
	return msg
}
