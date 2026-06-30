// internal/bot/keyboards/emergency.go
package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func EmergencyInline(formURL string) *tgbotapi.InlineKeyboardMarkup {
	if formURL == "" {
		return nil
	}

	btn := tgbotapi.NewInlineKeyboardButtonURL("🆘 Оставить заявку", formURL)
	row := tgbotapi.NewInlineKeyboardRow(btn)
	kb := tgbotapi.NewInlineKeyboardMarkup(row)
	return &kb
}
