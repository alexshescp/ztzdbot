// internal/bot/session/model.go
package session

type Step int

const (
	StepIdle Step = iota
	StepIncidentType
	StepIncidentDescription
	StepContact
)

type Session struct {
	ChatID        int64
	Step          Step
	IncidentType  string
	IncidentText  string
	ContactText   string
	LastMessageID int
}

func (s *Session) Reset() {
	s.Step = StepIdle
	s.IncidentType = ""
	s.IncidentText = ""
	s.ContactText = ""
	s.LastMessageID = 0
}
