package models

import "time"

type Log struct {
	ID        string    `json:"ID"`
	TimeStamp time.Time `json:"TimeStamp"`
	Log       string    `json:"Log"`

	// === RELATIONS ===

	SessionID string `json:"SessionID"`
	FlowID    string `json:"FlowID"`

	Flow    Flow    `gorm:"foreignKey:FlowID;references:Name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Session Session `gorm:"foreignKey:SessionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
