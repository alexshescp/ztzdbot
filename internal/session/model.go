// internal/bot/session/model.go
package session

import "t0t0dcyberbot/internal/domain"

type Step int

const (
	StepIdle Step = iota
	StepIncidentType
	StepIncidentDescription
	StepContact
)

type Flow string

const (
	FlowIncident  Flow = "incident"
	FlowEmergency Flow = "emergency"
)

type Session struct {
	ChatID       int64
	Flow         Flow
	Step         Step
	IncidentType domain.IncidentType
	IncidentText string
	ContactText  string
}

func (s *Session) Reset() {
	s.Flow = ""
	s.Step = StepIdle
	s.IncidentType = domain.IncidentTypeUnknown
	s.IncidentText = ""
	s.ContactText = ""
}
