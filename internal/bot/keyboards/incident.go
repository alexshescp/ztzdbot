// internal/bot/keyboards/incident.go
package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func IncidentTypes() *tgbotapi.InlineKeyboardMarkup {
	serverBtn := tgbotapi.NewInlineKeyboardButtonData("🖥 Сервер заражён", "incident_server")
	websiteBtn := tgbotapi.NewInlineKeyboardButtonData("🌐 Сайт взломан", "incident_website")
	ransomBtn := tgbotapi.NewInlineKeyboardButtonData("💀 Ransomware / шифровальщик", "incident_ransomware")

	row1 := tgbotapi.NewInlineKeyboardRow(serverBtn)
	row2 := tgbotapi.NewInlineKeyboardRow(websiteBtn)
	row3 := tgbotapi.NewInlineKeyboardRow(ransomBtn)

	kb := tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3)
	return &kb
}
