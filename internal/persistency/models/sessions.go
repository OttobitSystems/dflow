package models

import "time"

type Session struct {
	ID          string
	StartedAt   time.Time
	CompletedAt time.Time

	// === RELATIONS ===
	FlowID string

	Flow Flow `gorm:"foreignKey:FlowID;references:Name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Logs []Log `gorm:"foreignKey:SessionID;"`
}
