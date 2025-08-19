// Package repository: is all functionality to work with database
package repository

import (
	"bytes"
	"dflow/internal/persistency/models"
	"encoding/json"
	"log"
	"net/http"
)

var apiGatewayURL = "https://h6v6viz9bh.execute-api.eu-south-1.amazonaws.com/persist-data"

func persistData(cognitoToken string, Flows []models.Flow, Sessions []models.Session, Logs []models.Log) {
	requestParameters := PersistDataBody{
		ClientID: "OBS2025",
		SpaceID:  "Space01",
		Flows:    Flows,
		Sessions: Sessions,
		Logs:     Logs,
	}

	requestBody, _ := json.Marshal(requestParameters)

	req, err := http.NewRequest("POST", apiGatewayURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+cognitoToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error on request: %v", err)
	}
	defer resp.Body.Close()
}

type PersistDataBody struct {
	ClientID string           `json:"ClientID"`
	SpaceID  string           `json:"SpaceID"`
	Flows    []models.Flow    `json:"Flows"`
	Sessions []models.Session `json:"Sessions"`
	Logs     []models.Log     `json:"Logs"`
}
