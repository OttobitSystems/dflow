package models

import "time"

type Log struct {
	ID        string
	TimeStamp time.Time
	Log       string

	// === RELATIONS ===

	SessionID string
	FlowID    string

	Flow    Flow    `gorm:"foreignKey:FlowID;references:Name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Session Session `gorm:"foreignKey:SessionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
