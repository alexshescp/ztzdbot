// internal/domain/incident_types.go
package domain

type IncidentType string

const (
	IncidentTypeServer     IncidentType = "server"
	IncidentTypeWebsite    IncidentType = "website"
	IncidentTypeRansomware IncidentType = "ransomware"
	IncidentTypeUnknown    IncidentType = "unknown"
)
