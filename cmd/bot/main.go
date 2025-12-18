package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"t0t0dcyberbot/internal/domain"
	"t0t0dcyberbot/internal/infrastructure"
	"t0t0dcyberbot/internal/services/emergency"
	"t0t0dcyberbot/internal/services/messaging"
	"t0t0dcyberbot/internal/session"
)

const defaultAdminChatID int64 = 943354460

type botApp struct {
	bot         *tgbotapi.BotAPI
	store       *session.Store
	messaging   *messaging.Service
	emergency   *emergency.Service
	adminChatID int64
}

func main() {
	_ = godotenv.Load()

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is required")
	}

	adminChatID := parseAdminChatID(os.Getenv("ADMIN_CHAT_ID"))
	emergencyAPIURL := os.Getenv("EMERGENCY_API_URL")

	emergencyAPI := infrastructure.NewEmergencyAPI(emergencyAPIURL, nil)
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("failed to create bot: %v", err)
	}

	app := &botApp{
		bot:         bot,
		store:       session.NewStore(),
		messaging:   messaging.NewService(),
		emergency:   emergency.NewService(emergencyAPI),
		adminChatID: adminChatID,
	}

	app.run()
}

func parseAdminChatID(raw string) int64 {
	if raw == "" {
		return defaultAdminChatID
	}

	id, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		log.Printf("invalid ADMIN_CHAT_ID %q, using default %d", raw, defaultAdminChatID)
		return defaultAdminChatID
	}

	return id
}

func (a *botApp) run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := a.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.Text == "" {
			continue
		}
		a.handleMessage(update.Message)
	}
}

func (a *botApp) handleMessage(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	sess := a.store.Get(chatID)
	text := strings.TrimSpace(msg.Text)

	if handled := a.handleCommand(sess, chatID, text); handled {
		return
	}

	switch sess.Step {
	case session.StepIncidentType:
		sess.IncidentType = inferIncidentType(text)
		sess.Flow = session.FlowIncident
		sess.Step = session.StepIncidentDescription
		a.sendMarkdown(chatID, a.messaging.FormatAskIncidentDescription(sess.IncidentType), nil)
	case session.StepIncidentDescription:
		sess.IncidentText = text
		if sess.IncidentType == domain.IncidentTypeUnknown {
			sess.IncidentType = inferIncidentType(text)
		}
		sess.Step = session.StepContact
		if sess.Flow == session.FlowEmergency {
			a.sendMarkdown(chatID, a.messaging.FormatEmergencyContactRequest(sess.IncidentType), nil)
			return
		}
		sess.Flow = session.FlowIncident
		a.sendMarkdown(chatID, a.messaging.FormatAskContact(sess.IncidentType), nil)
	case session.StepContact:
		sess.ContactText = text
		a.submitIncident(chatID, sess)
	default:
		sess.Reset()
		a.sendMarkdown(chatID, a.messaging.FormatWelcome(), mainKeyboard())
	}
}

func (a *botApp) handleCommand(sess *session.Session, chatID int64, text string) bool {
	switch text {
	case "/start":
		sess.Reset()
		a.sendMarkdown(chatID, a.messaging.FormatWelcome(), mainKeyboard())
		return true
	case "ℹ️ О боте":
		a.sendMarkdown(chatID, a.messaging.FormatAbout(), mainKeyboard())
		return true
	case "🧨 Сообщить об инциденте":
		sess.Reset()
		sess.Flow = session.FlowIncident
		sess.Step = session.StepIncidentType
		a.sendMarkdown(chatID, a.messaging.FormatIncidentStart(), mainKeyboard())
		return true
	case "🆘 Экстренная поддержка":
		sess.Reset()
		sess.Flow = session.FlowEmergency
		sess.Step = session.StepIncidentDescription
		a.sendMarkdown(chatID, a.messaging.FormatEmergencyIntro(), mainKeyboard())
		return true
	default:
		return false
	}
}

func (a *botApp) submitIncident(chatID int64, sess *session.Session) {
	if sess.IncidentType == "" {
		sess.IncidentType = domain.IncidentTypeUnknown
	}
	if sess.Flow == "" {
		sess.Flow = session.FlowIncident
	}

	if err := a.emergency.SubmitIncident(sess.ContactText, sess.IncidentType, sess.IncidentText); err != nil {
		log.Printf("failed to submit incident: %v", err)
		a.sendMarkdown(chatID, "Не удалось отправить заявку, попробуйте ещё раз позже.", mainKeyboard())
		return
	}

	a.notifyAdmin(chatID, sess.IncidentType, sess.IncidentText, sess.ContactText, sess.Flow == session.FlowEmergency)

	if sess.Flow == session.FlowEmergency {
		a.sendMarkdown(chatID, a.messaging.FormatEmergencyReceived(sess.IncidentType, sess.IncidentText, sess.ContactText), mainKeyboard())
	} else {
		a.sendMarkdown(chatID, a.messaging.FormatIncidentReceived(sess.IncidentType, sess.IncidentText, sess.ContactText), mainKeyboard())
	}

	sess.Reset()
}

func (a *botApp) notifyAdmin(userChatID int64, incidentType domain.IncidentType, incidentText, contact string, emergency bool) {
	if a.adminChatID == 0 {
		return
	}
	adminText := a.messaging.FormatAdminNotification(userChatID, incidentType, incidentText, contact, emergency)
	a.sendMarkdown(a.adminChatID, adminText, nil)
}

func (a *botApp) sendMarkdown(chatID int64, text string, markup interface{}) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "Markdown"
	if markup != nil {
		msg.ReplyMarkup = markup
	}
	if _, err := a.bot.Send(msg); err != nil {
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

func inferIncidentType(text string) domain.IncidentType {
	result := domain.AnalyzeIncident(domain.AnalysisInput{Text: text})
	return result.GuessedType
}
