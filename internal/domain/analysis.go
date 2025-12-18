// internal/domain/analysis.go
package domain

import "strings"

type AnalysisInput struct {
	Text string
}

type AnalysisResult struct {
	GuessedType IncidentType
	Confidence  float64
}

func AnalyzeIncident(input AnalysisInput) AnalysisResult {
	t := strings.ToLower(input.Text)

	serverKeywords := []string{
		"сервер", "server", "cpu", "нагруз", "процесс", "xmrig", "майнер", "ssh", "rdp", "подключени",
	}
	websiteKeywords := []string{
		"сайт", "redirect", "редирект", "wordpress", "joomla", "битрикс", "bitrix", "deface", "hacked", "домен", "url",
	}
	ransomwareKeywords := []string{
		"шифроваль", "шифр", "ransom", "выкуп", ".lock", ".encrypted", "decrypt", "ransomware", "расшифров", "крипто",
	}

	for _, kw := range ransomwareKeywords {
		if strings.Contains(t, kw) {
			return AnalysisResult{GuessedType: IncidentTypeRansomware, Confidence: 0.9}
		}
	}

	for _, kw := range websiteKeywords {
		if strings.Contains(t, kw) {
			return AnalysisResult{GuessedType: IncidentTypeWebsite, Confidence: 0.8}
		}
	}

	for _, kw := range serverKeywords {
		if strings.Contains(t, kw) {
			return AnalysisResult{GuessedType: IncidentTypeServer, Confidence: 0.7}
		}
	}

	return AnalysisResult{GuessedType: IncidentTypeUnknown, Confidence: 0.3}
}
