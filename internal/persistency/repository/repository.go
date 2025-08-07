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

	_ = DBInstance.Find(&flows, `Name = "default"`)

	if len(flows) == 0 {
		CreateFlow("default")
		fmt.Println("`default` flow created!")
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

func InitSession(flowName string) (string, error) {
	newSection := models.Session{
		Id:     uuid.New().String(),
		FlowId: flowName,
	}

	var flows []models.Flow

	_ = DBInstance.First(&flows, `Name = "`+flowName+`"`)

	if len(flows) == 0 {
		return "", fmt.Errorf("FlowId not found")
	}

	_ = DBInstance.Create(&newSection)

	return newSection.Id, nil
}

func NotifySessionStarted(InDatabaseID string, StartedAt time.Time) error {
	var sessions []models.Session

	_ = DBInstance.First(&sessions, `Id = "`+InDatabaseID+`"`)

	if len(sessions) == 0 {
		return fmt.Errorf("session id not found")
	}

	session := sessions[0]

	session.StartedAt = StartedAt

	DBInstance.Model(&session).Where(`Id = "`+InDatabaseID+`"`).Update("StartedAt", StartedAt)

	return nil
}
