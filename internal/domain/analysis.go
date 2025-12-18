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

	if strings.Contains(t, "шифроваль") ||
		strings.Contains(t, "ransom") ||
		strings.Contains(t, "выкуп") ||
		strings.Contains(t, ".lock") ||
		strings.Contains(t, ".encrypted") {
		return AnalysisResult{GuessedType: IncidentTypeRansomware, Confidence: 0.9}
	}

	if strings.Contains(t, "редирект") ||
		strings.Contains(t, "redirect") ||
		strings.Contains(t, "wordpress") ||
		strings.Contains(t, "joomla") ||
		strings.Contains(t, "битрикс") ||
		strings.Contains(t, "bitrix") ||
		strings.Contains(t, "hacked by") {
		return AnalysisResult{GuessedType: IncidentTypeWebsite, Confidence: 0.8}
	}

	if strings.Contains(t, "cpu") ||
		strings.Contains(t, "нагруз") ||
		strings.Contains(t, "процесс") ||
		strings.Contains(t, "xmrig") ||
		strings.Contains(t, "майнер") ||
		strings.Contains(t, "ssh") {
		return AnalysisResult{GuessedType: IncidentTypeServer, Confidence: 0.7}
	}

	return AnalysisResult{GuessedType: IncidentTypeUnknown, Confidence: 0.3}
}
