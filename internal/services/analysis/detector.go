package analysis

import "strings"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) DetectType(text string) string {
	t := strings.ToLower(text)

	if strings.Contains(t, "шифров") ||
		strings.Contains(t, "ransom") ||
		strings.Contains(t, ".encrypted") {
		return "ransomware"
	}

	if strings.Contains(t, "redirect") ||
		strings.Contains(t, "wordpress") ||
		strings.Contains(t, "hacked") {
		return "website"
	}

	if strings.Contains(t, "cpu") ||
		strings.Contains(t, "майн") ||
		strings.Contains(t, "xmrig") ||
		strings.Contains(t, "ssh") {
		return "server"
	}

	return "unknown"
}
