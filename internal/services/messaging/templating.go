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
	b.WriteString("• аккуратно подготовить заявку для инженера 0t0d.\n\n")
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
	b.WriteString("\nПосле выбора типа я задам пару уточняющих вопросов.\n")
	return b.String()
}

func (s *Service) FormatAskIncidentDescription(incidentType string) string {
	it := domain.IncidentType(incidentType)
	short := domain.ShortTypeLabel(it)

	var b strings.Builder
	b.WriteString(fmt.Sprintf("*Опишите, что происходит (%s)*.\n\n", short))
	b.WriteString("Полезно указать:\n")
	b.WriteString("• какие именно симптомы вы видите (ошибки, сообщения, надписи);\n")
	b.WriteString("• когда примерено это началось;\n")
	b.WriteString("• что меняли/обновляли незадолго до инцидента;\n")
	b.WriteString("• есть ли скриншоты / логи.\n\n")
	b.WriteString("Напишите описание в одном-двух сообщениях. Потом я попрошу контакт для связи с инженером.\n")
	return b.String()
}

func (s *Service) FormatAskContact(incidentType, incidentText string) string {
	it := domain.IncidentType(incidentType)
	short := domain.ShortTypeLabel(it)

	var b strings.Builder
	b.WriteString("*Спасибо, зафиксировал описание инцидента.*\n\n")
	b.WriteString(fmt.Sprintf("Тип: _%s_.\n\n", short))
	b.WriteString("Теперь, пожалуйста, оставьте контакт для связи:\n")
	b.WriteString("• Telegram @username,\n")
	b.WriteString("• или номер телефона,\n")
	b.WriteString("• или email.\n\n")
	b.WriteString("Инженер 0t0d свяжется с вами, чтобы пройтись по шагам и помочь локализовать проблему.\n")
	return b.String()
}

func (s *Service) FormatIncidentReceived(incidentType, incidentText, contact string) string {
	it := domain.IncidentType(incidentType)
	short := domain.ShortTypeLabel(it)

	var b strings.Builder
	b.WriteString("*Заявка по инциденту передана инженеру 0t0d.*\n\n")
	b.WriteString(fmt.Sprintf("Тип: _%s_.\n", short))
	b.WriteString(fmt.Sprintf("Контакт: `%s`.\n\n", contact))
	b.WriteString("*Кратко по описанию:*\n")
	b.WriteString("```")
	if len(incidentText) > 500 {
		b.WriteString(incidentText[:500] + "...")
	} else {
		b.WriteString(incidentText)
	}
	b.WriteString("```\n\n")
	b.WriteString(domain.UniversalFiveSteps())
	b.WriteString("\nПока инженер подключается, пожалуйста, по возможности ограничьте доступ к системе и не удаляйте подозрительные файлы и логи.\n")
	return b.String()
}

func (s *Service) FormatEmergencyIntro() string {
	var b strings.Builder
	b.WriteString("*Экстренная поддержка 0t0d* 🆘\n\n")
	b.WriteString("Если ситуация критичная (простои, шифрование данных, массовые ошибки, требования выкупа) — лучше сразу оставить контакт для инженера.\n\n")
	b.WriteString("Напишите в одном сообщении:\n")
	b.WriteString("• ваш Telegram @username, телефон или email;\n")
	b.WriteString("• кратко: что именно сломалось (сервер, сайт, рабочие станции, база данных и т.п.).\n\n")
	b.WriteString("Я сформирую заявку и передам её на срочную обработку.\n")
	return b.String()
}

func (s *Service) FormatEmergencyWithForm() string {
	var b strings.Builder
	b.WriteString("*Экстренная заявка открыта.*\n\n")
	b.WriteString("Вы можете заполнить форму по кнопке выше или прислать сюда:\n")
	b.WriteString("• контакт для связи (Telegram / телефон / email);\n")
	b.WriteString("• два-три предложения, что именно произошло.\n\n")
	b.WriteString("Далее инженер 0t0d свяжется с вами напрямую.\n")
	return b.String()
}
