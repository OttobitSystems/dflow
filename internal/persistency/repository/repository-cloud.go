// Package repository: is all functionality to work with database
package repository

import (
	"bytes"
	"dflow/internal/persistency/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func persistData(cognitoToken string, Flows []models.Flow, Sessions []models.Session, Logs []models.Log) {
	apiGatewayURL := "https://h6v6viz9bh.execute-api.eu-south-1.amazonaws.com/persist-data"
	requestParameters := persistDataBody{
		ClientID: ApplicationConfiguration.ClientID,
		SpaceID:  ApplicationConfiguration.JoinedSpace,
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

func JoinCloudSpace(cognitoToken string, SpaceID string, ClientID string) bool {
	apiGatewayURL := "https://h6v6viz9bh.execute-api.eu-south-1.amazonaws.com/join-space"
	requestParameters := joinSpace{
		ClientID: ClientID,
		SpaceID:  SpaceID,
	}

	requestBody, _ := json.Marshal(requestParameters)

	req, err := http.NewRequest("POST", apiGatewayURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return false
	}

	req.Header.Set("Authorization", "Bearer "+cognitoToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error on request: %v", err)
		return false
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading body: %v", err)
		return false
	}

	var joinResponse joinResponse
	err = json.Unmarshal(body, &joinResponse)
	if err != nil {
		log.Fatalf("error parsing body: %v", err)
		return false
	}

	return joinResponse.ClientExists
}

func AddCloudSpace(cognitoToken string, SpaceID string, ClientID string) string {
	apiGatewayURL := "https://h6v6viz9bh.execute-api.eu-south-1.amazonaws.com/space-add"
	requestParameters := createSpace{
		ClientID: ClientID,
		SpaceID:  SpaceID,
	}

	requestBody, _ := json.Marshal(requestParameters)

	req, err := http.NewRequest("POST", apiGatewayURL, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+cognitoToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("error on request: %v", err)
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error reading body: %v", err)
		panic(err)
	}

	var createSpaceResponse createSpaceResponse
	err = json.Unmarshal(body, &createSpaceResponse)
	if err != nil {
		log.Fatalf("error parsing body: %v", err)
		panic(err)
	}

	return createSpaceResponse.Message
}

type persistDataBody struct {
	ClientID string           `json:"ClientID"`
	SpaceID  string           `json:"SpaceID"`
	Flows    []models.Flow    `json:"Flows"`
	Sessions []models.Session `json:"Sessions"`
	Logs     []models.Log     `json:"Logs"`
}

type joinSpace struct {
	ClientID string `json:"ClientID"`
	SpaceID  string `json:"SpaceID"`
}

type joinResponse struct {
	ClientExists bool
}

type createSpace struct {
	ClientID string `json:"ClientID"`
	SpaceID  string `json:"SpaceID"`
}

type createSpaceResponse struct {
	Message string
}
