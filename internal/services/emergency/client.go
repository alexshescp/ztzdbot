package emergency

import (
	"fmt"

	"t0t0dcyberbot/internal/domain"
	"t0t0dcyberbot/internal/infrastructure"
)

type Service struct {
	api *infrastructure.EmergencyAPI
}

func NewService(api *infrastructure.EmergencyAPI) *Service {
	if api == nil {
		api = infrastructure.NewEmergencyAPI("", nil)
	}
	return &Service{api: api}
}

func (s *Service) SubmitIncident(contact string, incidentType domain.IncidentType, description string) error {
	payload := fmt.Sprintf("[%s]\n%s", incidentType, description)
	return s.api.Submit(contact, payload)
}
