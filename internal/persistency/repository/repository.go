package repository

import (
	"dflow/internal/persistency/models"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

var DbInstance *gorm.DB = nil

const SQLCONNECTIONSTRING = "dflow.db?cache=shared&foreign_keys=1"

func InitDatabase() (bool, error) {
	var err error
	DbInstance, err = gorm.Open(sqlite.Open(SQLCONNECTIONSTRING), &gorm.Config{})

	DbInstance.AutoMigrate(&models.Flow{}, &models.Session{}, &models.Log{})

	return (err != nil), err
}

func CreateFlow(name string) (bool, error) {

	newFlow := models.Flow{
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}

	status := DbInstance.Create(&newFlow)

	return (status != nil), status.Error
}

func StartSession(flowName string) (string, error) {
	newSection := models.Session{
		Id:        uuid.New().String(),
		FlowId:    flowName,
		StartedAt: time.Now().UTC(),
	}

	status := DbInstance.Create(&newSection)

	return newSection.Id, status.Error
}
