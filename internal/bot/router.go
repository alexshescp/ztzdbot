// internal/bot/router.go
package bot

import (
	"t0t0dcyberbot/internal/bot/handlers"
	"t0t0dcyberbot/internal/session"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Router struct {
	sessions *session.Store
	handlers *handlers.HandlerSet
}

func NewRouter(s *session.Store, h *handlers.HandlerSet) *Router {
	return &Router{
		sessions: s,
		handlers: h,
	}
}

func (r *Router) Handle(update tgbotapi.Update) tgbotapi.Chattable {
	if update.CallbackQuery != nil {
		return r.handleCallback(update)
	}
	if update.Message != nil {
		return r.handleMessage(update)
	}
	return nil
}

func (r *Router) handleMessage(update tgbotapi.Update) tgbotapi.Chattable {
	msg := update.Message
	chatID := msg.Chat.ID
	sess := r.sessions.Get(chatID)

	text := msg.Text

	switch text {
	case "start":
		return r.handlers.Start.HandleStart(update, sess)

	case "ℹ️ О боте":
		return r.handlers.Start.HandleAbout(update, sess)

	case "🧨 Сообщить об инциденте":
		return r.handlers.Incidents.HandleStartIncident(update, sess)

	case "🆘 Экстренная поддержка":
		return r.handlers.Emergency.HandleEmergency(update, sess)

	default:
		// если есть активный шаг инцидента — продолжаем сценарий
		if sess.Step != session.StepIdle {
			return r.handlers.Incidents.HandleStep(update, sess)
		}
		// иначе — fallback
		return r.handlers.Fallback.HandleUnknown(update, sess)
	}
}

func (r *Router) handleCallback(update tgbotapi.Update) tgbotapi.Chattable {
	cb := update.CallbackQuery
	chatID := cb.Message.Chat.ID
	sess := r.sessions.Get(chatID)

	switch cb.Data {
	case "incident_server":
		return r.handlers.Incidents.HandleSelectType(update, sess, "server")

	case "incident_website":
		return r.handlers.Incidents.HandleSelectType(update, sess, "website")

	case "incident_ransomware":
		return r.handlers.Incidents.HandleSelectType(update, sess, "ransomware")

	case "emergency_open_form":
		return r.handlers.Emergency.HandleOpenForm(update, sess)

	default:
		return r.handlers.Fallback.HandleUnknownCallback(update, sess)
	}
}
