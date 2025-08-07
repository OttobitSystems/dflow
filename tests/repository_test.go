package main

import (
	"dflow/internal/persistency/repository"
	"testing"
	"time"
)

func TestShouldGenerateNewDatabase(t *testing.T) {
	// Arrange
	expectedResult := true

	// Act
	databaseInitializationStatus, err := repository.InitDatabase()

	// Assert
	if databaseInitializationStatus != expectedResult || err != nil {
		t.Error("Error on init, error code: ", err)
	}
}

func TestShouldAddFlowInDatabase(t *testing.T) {
	// Arrange
	expectedResult := true

	// Act
	repository.InitDatabase()
	functionResult, err := repository.CreateFlow(time.Now().UTC().GoString())

	// Assert
	if functionResult != expectedResult {
		t.Error("error creating flow, ", err)
	}
}

func TestShouldAddSessionInDatabase(t *testing.T) {
	// Arrange

	// Act
	repository.InitDatabase()
	sessionId, err := repository.StartSession("TestFlow")

	// Assert
	if sessionId != "" {
		t.Error("errror creating session, err: ", err)
	}
}
