package pagerduty

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/mayuka-c/pagerduty-auto/pkg/pagerduty/models"
)

const (
	pdEndpoint = "https://api.pagerduty.com"
)

type errorStruct struct {
	ErrorString interface{} `json:"error"`
	Description string      `jsom:"error_description"`
}

type pdClient struct {
	token      string
	httpClient *resty.Client
}

func NewPDclient(token string) *pdClient {
	return &pdClient{httpClient: resty.New(), token: token}
}

func (p *pdClient) ListIncidents(ctx context.Context, userid string) ([]models.Incident, error) {
	result := models.PDIncidents{}
	errorS := errorStruct{}
	listIncidentsEndpoint := pdEndpoint + "/incidents"

	resp, err := p.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Token token="+p.token).
		SetQueryParams(map[string]string{"user_ids[]": userid}).
		SetResult(&result).SetError(&errorS).Get(listIncidentsEndpoint)
	if err != nil {
		return result.Incident, err
	}

	if resp.IsError() || !resp.IsSuccess() {
		return result.Incident, fmt.Errorf("got status code: %d with error: %v", resp.StatusCode(), errorS.ErrorString)
	}

	fmt.Println("Successfully got the list of incidents")
	return result.Incident, nil
}

func (p *pdClient) UpdateIncident(ctx context.Context, id string, req models.UpdateIncidentRequest) error {
	result := models.PDIncidents{}
	errorS := errorStruct{}
	updateEndpoint := pdEndpoint + "/incidents/" + id

	resp, err := p.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("Authorization", "Token token="+p.token).
		SetHeader("From", req.From).
		SetBody(map[string]interface{}{
			"incident": map[string]interface{}{
				"type":   req.IncidentType,
				"status": req.Status,
			},
		}).SetResult(&result).SetError(&errorS).Put(updateEndpoint)
	if err != nil {
		return err
	}

	if resp.IsError() || !resp.IsSuccess() {
		return fmt.Errorf("got status code: %d with error: %v", resp.StatusCode(), errorS.ErrorString)
	}

	fmt.Printf("Successfully updated incident with id: %s\n", id)
	return nil
}
