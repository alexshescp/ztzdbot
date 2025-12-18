package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"t0t0dcyberbot/internal/domain"
	"t0t0dcyberbot/internal/infrastructure"
	"t0t0dcyberbot/internal/services/analysis"
	"t0t0dcyberbot/internal/services/emergency"
	"t0t0dcyberbot/internal/services/messaging"
	"t0t0dcyberbot/internal/session"
)

func main() {
	_ = godotenv.Load()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is required")
	}

	emergencyAPI := infrastructure.NewEmergencyAPI(os.Getenv("EMERGENCY_API_URL"), nil)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("failed to create bot: %v", err)
	}

	store := session.NewStore()
	messagingService := messaging.NewService()
	analysisService := analysis.NewService()
	emergencyService := emergency.NewService(emergencyAPI)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.Text == "" {
			continue
		}

		chatID := update.Message.Chat.ID
		sess := store.Get(chatID)
		text := strings.TrimSpace(update.Message.Text)

		switch text {
		case "/start":
			sess.Reset()
			sendMarkdown(bot, chatID, messagingService.FormatWelcome(), mainKeyboard())
			continue
		case "ℹ️ О боте":
			sendMarkdown(bot, chatID, messagingService.FormatAbout(), mainKeyboard())
			continue
		case "🧨 Сообщить об инциденте":
			sess.Reset()
			sess.Step = session.StepIncidentType
			sendMarkdown(bot, chatID, messagingService.FormatIncidentStart(), mainKeyboard())
			continue
		case "🆘 Экстренная поддержка":
			sess.Reset()
			sess.Step = session.StepIncidentDescription
			sendMarkdown(bot, chatID, messagingService.FormatEmergencyIntro()+"\n\nНапишите кратко, что произошло. После этого попрошу контакт для связи.", mainKeyboard())
			continue
		}

		switch sess.Step {
		case session.StepIncidentType:
			incidentType := inferIncidentType(analysisService, text)
			sess.IncidentType = string(incidentType)
			sess.Step = session.StepIncidentDescription
			sendMarkdown(bot, chatID, messagingService.FormatAskIncidentDescription(string(incidentType)), nil)
		case session.StepIncidentDescription:
			sess.IncidentText = text
			sess.Step = session.StepContact
			sendMarkdown(bot, chatID, messagingService.FormatAskContact(sess.IncidentType, sess.IncidentText), nil)
		case session.StepContact:
			sess.ContactText = text
			if err := emergencyService.SubmitIncident(sess.ContactText, domain.IncidentType(sess.IncidentType), sess.IncidentText); err != nil {
				log.Printf("failed to submit incident: %v", err)
				sendMarkdown(bot, chatID, "Не удалось отправить заявку, попробуйте ещё раз позже.", mainKeyboard())
				continue
			}

			sendMarkdown(bot, chatID, messagingService.FormatIncidentReceived(sess.IncidentType, sess.IncidentText, sess.ContactText), mainKeyboard())
			sess.Reset()
		default:
			sendMarkdown(bot, chatID, messagingService.FormatWelcome(), mainKeyboard())
		}
	}
}

func sendMarkdown(bot *tgbotapi.BotAPI, chatID int64, text string, markup interface{}) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"
	if markup != nil {
		msg.ReplyMarkup = markup
	}
	if _, err := bot.Send(msg); err != nil {
		log.Printf("failed to send message: %v", err)
	}
}

func mainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🧨 Сообщить об инциденте"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("🆘 Экстренная поддержка"),
			tgbotapi.NewKeyboardButton("ℹ️ О боте"),
		),
	)
}

func inferIncidentType(service *analysis.Service, text string) domain.IncidentType {
	normalized := strings.ToLower(text)
	switch {
	case strings.Contains(normalized, "сервер"):
		return domain.IncidentTypeServer
	case strings.Contains(normalized, "сайт"):
		return domain.IncidentTypeWebsite
	case strings.Contains(normalized, "ransom"), strings.Contains(normalized, "шифров"):
		return domain.IncidentTypeRansomware
	}

	switch service.DetectType(text) {
	case "server":
		return domain.IncidentTypeServer
	case "website":
		return domain.IncidentTypeWebsite
	case "ransomware":
		return domain.IncidentTypeRansomware
	default:
		return domain.IncidentTypeUnknown
	}
}
