package handlers

type MessagingService interface {
    FormatWelcome() string
    FormatAbout() string
    FormatIncidentStart() string
    FormatAskIncidentDescription(incidentType string) string
    FormatAskContact(incidentType, incidentText string) string
    FormatIncidentReceived(incidentType, incidentText, contact string) string
    FormatEmergencyIntro() string
    FormatEmergencyWithForm() string
}

type EmergencyClient interface {
    FormURL() string
    SubmitRequest(contact, incident string) error
}
