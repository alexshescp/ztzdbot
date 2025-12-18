// internal/infrastructure/emergency_api.go
package infrastructure

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type EmergencyAPI struct {
	BaseURL string
	Client  HTTPClient
}

type EmergencyRequest struct {
	Contact string `json:"contact"`
	Text    string `json:"text"`
}

func NewEmergencyAPI(baseURL string, client HTTPClient) *EmergencyAPI {
	if client == nil {
		client = NewDefaultHTTPClient()
	}
	return &EmergencyAPI{
		BaseURL: baseURL,
		Client:  client,
	}
}

func (api *EmergencyAPI) Submit(contact, text string) error {
	if api.BaseURL == "" {
		return nil
	}

	body, err := json.Marshal(EmergencyRequest{
		Contact: contact,
		Text:    text,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, api.BaseURL, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	_, err = api.Client.Do(req)
	return err
}
