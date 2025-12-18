// internal/domain/instructions.go
package domain

import "strings"

func UniversalFiveSteps() string {
	var b strings.Builder
	b.WriteString("*Универсальный экстренный протокол (5 шагов)*\n\n")
	b.WriteString("1️⃣ *Изоляция*: временно ограничьте доступ к системе/сайту (файрвол, выключение внешнего доступа, пауза сервисов).\n")
	b.WriteString("2️⃣ *Сохранение артефактов*: не удаляйте странные файлы и заметки, сохраните логи и снимки экрана.\n")
	b.WriteString("3️⃣ *Диагностика*: зафиксируйте, что именно не работает, какие есть ошибки, когда это началось.\n")
	b.WriteString("4️⃣ *Минимизация ущерба*: если есть критические сервисы, временно переведите их на резерв или отдельный хост.\n")
	b.WriteString("5️⃣ *Подключение инженера*: дайте контакт для связи (Telegram / телефон / email), чтобы мы подключили специалиста.\n")
	return b.String()
}

func ShortTypeLabel(t IncidentType) string {
	switch t {
	case IncidentTypeServer:
		return "заражение сервера"
	case IncidentTypeWebsite:
		return "взлом сайта"
	case IncidentTypeRansomware:
		return "шифровальщик / ransomware"
	default:
		return "инцидент (тип уточняется)"
	}
}
