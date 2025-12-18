package messaging

import (
	"fmt"
	"strings"

	"t0t0dcyberbot/internal/domain"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) FormatWelcome() string {
	var b strings.Builder
	b.WriteString("*t0t0dcyberbot* — антикризисный кибербот по инцидентам безопасности 🚨\n\n")
	b.WriteString("Я помогу:\n")
	b.WriteString("• быстро сориентироваться по типу инцидента (сервер, сайт, ransomware);\n")
	b.WriteString("• дать понятный экстренный план в 5 шагов;\n")
	b.WriteString("• аккуратно подготовить заявку для инженера 0t0d и в админский чат.\n\n")
	b.WriteString("Выберите действие на клавиатуре ниже:\n")
	b.WriteString("• «🧨 Сообщить об инциденте» — если уже что-то случилось;\n")
	b.WriteString("• «🆘 Экстренная поддержка» — если нужно срочно подключить инженера;\n")
	b.WriteString("• «ℹ️ О боте» — чтобы понять, как мы работаем.\n")
	return b.String()
}

func (s *Service) FormatAbout() string {
	var b strings.Builder
	b.WriteString("*О t0t0dcyberbot и команде 0t0d*\n\n")
	b.WriteString("Мы занимаемся реагированием на инциденты:\n")
	b.WriteString("• заражение серверов и инфраструктуры;\n")
	b.WriteString("• взлом и дефейс сайтов, редиректы и фишинг;\n")
	b.WriteString("• шифровальщики / ransomware и вымогательство.\n\n")
	b.WriteString("Подход:\n")
	b.WriteString("1. Быстро уточняем симптомы и тип инцидента.\n")
	b.WriteString("2. Даём понятный экстренный план действий.\n")
	b.WriteString("3. Помогаем собрать минимальный набор артефактов для расследования.\n")
	b.WriteString("4. При необходимости подключаем инженера 0t0d для срочной диагностики.\n\n")
	b.WriteString("Вы можете в любой момент нажать «🧨 Сообщить об инциденте» или «🆘 Экстренная поддержка» на клавиатуре.\n")
	return b.String()
}

func (s *Service) FormatIncidentStart() string {
	var b strings.Builder
	b.WriteString("*Сообщить об инциденте безопасности*\n\n")
	b.WriteString("Для начала давайте определим, к какому типу ближе ситуация:\n\n")
	b.WriteString("• *Сервер заражён* — странные процессы, высокая нагрузка, майнинг, неизвестные подключения;\n")
	b.WriteString("• *Сайт взломан* — редиректы, «Hacked by ...», спам / фишинг, ломается верстка;\n")
	b.WriteString("• *Ransomware* — файлы зашифрованы, появилось требование выкупа.\n\n")
	b.WriteString(domain.UniversalFiveSteps())
	b.WriteString("\nПосле выбора типа я попрошу описать ситуацию и оставить контакт — заявка уйдёт инженеру 0t0d и в админский чат.\n")
	return b.String()
}

func (s *Service) FormatAskIncidentDescription(incidentType domain.IncidentType) string {
	short := domain.ShortTypeLabel(incidentType)

	var b strings.Builder
	b.WriteString(fmt.Sprintf("*Опишите, что происходит (%s).*\n\n", short))
	b.WriteString("Полезно указать:\n")
	b.WriteString("• какие симптомы вы видите (ошибки, надписи, странные файлы);\n")
	b.WriteString("• когда это началось и что меняли незадолго до инцидента;\n")
	b.WriteString("• есть ли скриншоты или логи.\n\n")
	b.WriteString("После описания я попрошу контакт для инженера.\n")
	return b.String()
}

func (s *Service) FormatAskContact(incidentType domain.IncidentType) string {
	short := domain.ShortTypeLabel(incidentType)

	var b strings.Builder
	b.WriteString("*Спасибо, зафиксировал описание инцидента.*\n\n")
	b.WriteString(fmt.Sprintf("Тип: _%s_.\n\n", short))
	b.WriteString("Оставьте контакт для связи:\n")
	b.WriteString("• Telegram @username,\n")
	b.WriteString("• или номер телефона,\n")
	b.WriteString("• или email.\n\n")
	b.WriteString("Заявка уйдёт инженеру 0t0d и в админский чат, чтобы ускорить реакцию.\n")
	return b.String()
}

func (s *Service) FormatIncidentReceived(incidentType domain.IncidentType, incidentText, contact string) string {
	short := domain.ShortTypeLabel(incidentType)

	var b strings.Builder
	b.WriteString("*Заявка по инциденту передана инженеру 0t0d.*\n\n")
	b.WriteString(fmt.Sprintf("Тип: _%s_.\n", short))
	b.WriteString(fmt.Sprintf("Контакт: `%s`.\n\n", sanitizeInline(contact)))
	b.WriteString("*Кратко по описанию:*\n")
	b.WriteString(formatCodeBlock(truncateText(incidentText, 800)))
	b.WriteString("\n")
	b.WriteString(domain.UniversalFiveSteps())
	b.WriteString("\nПока инженер подключается, пожалуйста, по возможности ограничьте доступ к системе и не удаляйте подозрительные файлы и логи.\n")
	return b.String()
}

func (s *Service) FormatEmergencyIntro() string {
	var b strings.Builder
	b.WriteString("*Экстренная поддержка 0t0d* 🆘\n\n")
	b.WriteString("Если ситуация критичная (простои, шифрование данных, массовые ошибки, требования выкупа) — давайте зафиксируем её сразу.\n\n")
	b.WriteString("Сначала опишите, что произошло. После этого попрошу контакт, чтобы подключить инженера напрямую.\n")
	return b.String()
}

func (s *Service) FormatEmergencyContactRequest(incidentType domain.IncidentType) string {
	short := domain.ShortTypeLabel(incidentType)

	var b strings.Builder
	b.WriteString("*Принял экстренное описание.*\n\n")
	b.WriteString(fmt.Sprintf("Предварительный тип: _%s_.\n\n", short))
	b.WriteString("Оставьте контакт для связи:\n")
	b.WriteString("• Telegram @username,\n")
	b.WriteString("• номер телефона,\n")
	b.WriteString("• или email.\n\n")
	b.WriteString("Передам заявку инженеру и в админский чат для ускоренной реакции.\n")
	return b.String()
}

func (s *Service) FormatEmergencyReceived(incidentType domain.IncidentType, incidentText, contact string) string {
	short := domain.ShortTypeLabel(incidentType)

	var b strings.Builder
	b.WriteString("*Экстренная заявка принята и передана инженеру 0t0d.*\n\n")
	b.WriteString(fmt.Sprintf("Тип: _%s_.\n", short))
	b.WriteString(fmt.Sprintf("Контакт: `%s`.\n\n", sanitizeInline(contact)))
	b.WriteString("*Кратко по описанию:*\n")
	b.WriteString(formatCodeBlock(truncateText(incidentText, 900)))
	b.WriteString("\nДержите связь — инженер напишет или позвонит, как только получит заявку.\n")
	return b.String()
}

func (s *Service) FormatAdminNotification(chatID int64, incidentType domain.IncidentType, incidentText, contact string, emergency bool) string {
	short := domain.ShortTypeLabel(incidentType)

	var b strings.Builder
	if emergency {
		b.WriteString("*Новая заявка (экстренная)* 🆘\n\n")
	} else {
		b.WriteString("*Новая заявка по инциденту* 🧨\n\n")
	}

	b.WriteString(fmt.Sprintf("Чат пользователя: `%d`\n", chatID))
	b.WriteString(fmt.Sprintf("Тип: %s\n", short))
	b.WriteString(fmt.Sprintf("Контакт: `%s`\n\n", sanitizeInline(contact)))
	b.WriteString("*Описание:*\n")
	b.WriteString(formatCodeBlock(truncateText(incidentText, 1000)))
	return b.String()
}

func formatCodeBlock(text string) string {
	return "```" + sanitizeForBlock(text) + "```"
}

func truncateText(text string, limit int) string {
	if len(text) <= limit {
		return text
	}
	return text[:limit] + "..."
}

func sanitizeInline(text string) string {
	text = strings.ReplaceAll(text, "`", "´")
	text = strings.ReplaceAll(text, "\n", " ")
	return strings.TrimSpace(text)
}

func sanitizeForBlock(text string) string {
	text = strings.ReplaceAll(text, "```", "'''")
	return text
}
