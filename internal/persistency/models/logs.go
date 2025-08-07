package models

import "time"

type Log struct {
	Id        string
	TimeStamp time.Time
	Log       string

	// === RELATIONS ===

	SessionId string
	FlowId    string

	Flow    Flow    `gorm:"foreignKey:FlowId;references:Name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Session Session `gorm:"foreignKey:SessionId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
