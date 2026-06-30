package handlers

import (
	"t0t0dcyberbot/internal/services/analysis"
	"t0t0dcyberbot/internal/services/emergency"
	"t0t0dcyberbot/internal/services/instructions"
	"t0t0dcyberbot/internal/services/messaging"
)

type HandlerSet struct {
	Start     *StartHandler
	Emergency *EmergencyHandler
	Incidents *IncidentsHandler
	Fallback  *FallbackHandler
}

func NewHandlerSet(
	analysisSvc *analysis.Service,
	instructionsSvc *instructions.Service,
	msgSvc *messaging.Service,
	emergencyClient *emergency.Client,
) *HandlerSet {

	start := &StartHandler{
		Msg: msgSvc,
	}

	emer := &EmergencyHandler{
		Msg:       msgSvc,
		Emergency: emergencyClient,
	}

	inc := &IncidentsHandler{
		Msg:       msgSvc,
		Emergency: emergencyClient,
	}

	fallback := &FallbackHandler{
		Msg: msgSvc,
	}

	return &HandlerSet{
		Start:     start,
		Emergency: emer,
		Incidents: inc,
		Fallback:  fallback,
	}
}
