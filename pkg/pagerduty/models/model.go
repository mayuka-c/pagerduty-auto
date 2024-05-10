package models

type PDIncidents struct {
	Incident []Incident `json:"incidents"`
}

type Incident struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type UpdateIncidentRequest struct {
	From         string
	IncidentType string
	Status       string
}
