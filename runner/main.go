package main

import (
	"context"
	"fmt"

	"github.com/mayuka-c/pagerduty-auto/pkg/pagerduty"
	"github.com/mayuka-c/pagerduty-auto/pkg/pagerduty/models"
	"github.com/mayuka-c/pagerduty-auto/pkg/parser"
)

var (
	ctx = context.Background()
)

func getTriggeredIncident(incidents []models.Incident) []models.Incident {
	filteredIncidents := []models.Incident{}
	for _, incident := range incidents {
		if incident.Status == "triggered" {
			filteredIncidents = append(filteredIncidents, incident)
		}
	}
	return filteredIncidents
}

func main() {
	parser.Parse()
	if parser.InputFlags.Email == "" || parser.InputFlags.ID == "" {
		fmt.Println("Please provide email and ID and Token")
		return
	}

	pdClient := pagerduty.NewPDclient(parser.InputFlags.Token)

	incidents, err := pdClient.ListIncidents(ctx, parser.InputFlags.ID)
	if err != nil {
		fmt.Printf("Failed to get incidents with err: %v\n", err)
		return
	}

	for _, incident := range getTriggeredIncident(incidents) {
		err := pdClient.UpdateIncident(ctx, incident.Id, models.UpdateIncidentRequest{
			From:         parser.InputFlags.Email,
			IncidentType: "incident_reference",
			Status:       "acknowledged",
		})
		if err != nil {
			fmt.Printf("Failed to update incident for id: %v with err: %v\n", incident.Id, err)
		}
	}

	fmt.Printf("All the incidents for this user: %v are successfully acknowledged\n", parser.InputFlags.Email)
}
