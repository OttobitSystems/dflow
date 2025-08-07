package repository

import (
	"dflow/internal/persistency/models"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBInstance *gorm.DB = nil

const SQLCONNECTIONSTRING = "dflow.db?cache=shared&foreign_keys=1"

func InitDatabase() (bool, error) {
	var err error
	result := false
	DBInstance, err = gorm.Open(sqlite.Open(SQLCONNECTIONSTRING), &gorm.Config{})

	DBInstance.AutoMigrate(&models.Flow{}, &models.Session{}, &models.Log{})

	var flows []models.Flow

	_ = DBInstance.Find(&flows, `Name = "Default"`)

	if len(flows) == 0 {
		CreateFlow("Default")
		fmt.Println("Default flow created!")
	}

	result = true

	return result, err
}

func CreateFlow(name string) (bool, error) {
	newFlow := models.Flow{
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}

	status := DBInstance.Create(&newFlow)

	return (status != nil), status.Error
}

func StartSession(flowName string) (string, error) {
	newSection := models.Session{
		Id:        uuid.New().String(),
		FlowId:    flowName,
		StartedAt: time.Now().UTC(),
	}

	status := DBInstance.Create(&newSection)

	return newSection.Id, status.Error
}
