// Package repository: is all functionality to work with database
package repository

import (
	"dflow/internal/cloud/auth"
	"dflow/internal/persistency/models"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBInstance               *gorm.DB = nil
	ApplicationConfiguration          = &models.ApplicationConfiguration{
		DefaultFlow: "default",
	}
)

const SQLCONNECTIONSTRING = "dflow.db?cache=shared&foreign_keys=1"

func InitDatabase() (bool, error) {
	var err error
	result := false
	DBInstance, err = gorm.Open(sqlite.Open(SQLCONNECTIONSTRING), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DBInstance.AutoMigrate(&models.Flow{}, &models.Session{}, &models.Log{}, &models.ApplicationConfiguration{})

	var flows []models.Flow
	var configurations []models.ApplicationConfiguration

	configResult := DBInstance.First(&configurations)

	if configResult.Error != nil {
		DBInstance.Create(ApplicationConfiguration)
	}

	if configResult.Error == nil {
		ApplicationConfiguration = &configurations[0]
	}

	_ = DBInstance.Find(&flows, `Name = "`+ApplicationConfiguration.DefaultFlow+`"`)

	if len(flows) == 0 {
		CreateFlow(ApplicationConfiguration.DefaultFlow)
		fmt.Println("`" + ApplicationConfiguration.DefaultFlow + "` flow created!")
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

	if auth.UserLogedInCloud {
		token := auth.RefreshSession()
		persistData(token, []models.Flow{newFlow}, nil, nil)
	}

	return (status != nil), status.Error
}

func GetFlows() []models.Flow {
	var flows []models.Flow
	_ = DBInstance.Find(&flows)

	return flows
}

func GetFlowsAndSessions() []models.Flow {
	var flows []models.Flow
	_ = DBInstance.Preload("Sessions").Find(&flows)

	return flows
}

func InitSession(flowName string) (string, error) {
	newSection := models.Session{
		ID:     uuid.New().String(),
		FlowID: flowName,
	}

	var flows []models.Flow

	_ = DBInstance.First(&flows, `Name = "`+flowName+`"`)

	if len(flows) == 0 {
		return "", fmt.Errorf("FlowId not found")
	}

	_ = DBInstance.Create(&newSection)

	return newSection.ID, nil
}

func NotifySessionStarted(InDatabaseID string, StartedAt time.Time) error {
	var sessions []models.Session

	_ = DBInstance.First(&sessions, `ID = "`+InDatabaseID+`"`)

	if len(sessions) == 0 {
		return fmt.Errorf("session id not found")
	}

	session := sessions[0]

	session.StartedAt = StartedAt

	DBInstance.Model(&session).Where(`ID = "`+InDatabaseID+`"`).Update("StartedAt", StartedAt)

	return nil
}

func NotifySessionEnd(InDatabaseID string, CompletedAt time.Time) error {
	var sessions []models.Session

	_ = DBInstance.First(&sessions, `ID = "`+InDatabaseID+`"`)

	if len(sessions) == 0 {
		return fmt.Errorf("session id not found")
	}

	session := sessions[0]

	session.CompletedAt = CompletedAt

	DBInstance.Model(&session).Where(`ID = "`+InDatabaseID+`"`).Update("CompletedAt", CompletedAt)

	if auth.UserLogedInCloud {
		token := auth.RefreshSession()
		persistData(token, nil, []models.Session{session}, nil)
	}

	return nil
}

func GetAllLastLogs(SessionID string, FlowID string) []models.Log {
	var logs []models.Log

	_ = DBInstance.Order("time_stamp desc").Limit(10).Find(&logs, &models.Log{FlowID: FlowID, SessionID: SessionID})

	return logs
}

func GetLogs(FlowID string) []models.Log {
	var logs []models.Log

	_ = DBInstance.Order("time_stamp desc").Preload("Session").Find(&logs, &models.Log{FlowID: FlowID})

	return logs
}

func StoreLog(SessionID string, FlowID string, logText string) error {
	messageToLog := models.Log{
		ID:        uuid.New().String(),
		FlowID:    FlowID,
		SessionID: SessionID,
		TimeStamp: time.Now(),
		Log:       logText,
	}

	DBInstance.Create(messageToLog)

	if auth.UserLogedInCloud {
		token := auth.RefreshSession()
		persistData(token, nil, nil, []models.Log{messageToLog})
	}

	return nil
}

func UpdateDefaultFlowName(DefaultFlowName string) {
	var configurations []models.ApplicationConfiguration
	DBInstance.Model(&configurations).Where(`default_flow = "`+ApplicationConfiguration.DefaultFlow+`"`).Update("DefaultFlow", DefaultFlowName)
}

func UpdateClientID(ClientID string) {
	var configurations []models.ApplicationConfiguration
	DBInstance.Model(&configurations).Where(`default_flow = "`+ApplicationConfiguration.DefaultFlow+`"`).Update("client_id", ClientID)
}

func UpdateSpaceID(SpaceID string) {
	var configurations []models.ApplicationConfiguration
	DBInstance.Model(&configurations).Where(`default_flow = "`+ApplicationConfiguration.DefaultFlow+`"`).Update("joined_space", SpaceID)
}
