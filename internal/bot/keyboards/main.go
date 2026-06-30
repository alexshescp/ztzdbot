// internal/bot/keyboards/main.go
package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Main() tgbotapi.ReplyKeyboardMarkup {
	row1 := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🧨 Сообщить об инциденте"),
	)
	row2 := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("ℹ️ О боте"),
		tgbotapi.NewKeyboardButton("🆘 Экстренная поддержка"),
	)

	kb := tgbotapi.NewReplyKeyboard(row1, row2)
	kb.ResizeKeyboard = true
	kb.OneTimeKeyboard = false
	return kb
}
