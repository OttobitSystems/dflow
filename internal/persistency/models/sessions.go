package models

import "time"

type Session struct {
	Id          string
	StartedAt   time.Time
	CompletedAt time.Time

	// === RELATIONS ===
	FlowId string

	Flow Flow `gorm:"foreignKey:FlowId;references:Name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Logs []Log `gorm:"foreignKey:SessionId;"`
}
