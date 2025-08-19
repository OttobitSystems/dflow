package models

import "time"

type Session struct {
	ID          string    `json:"ID"`
	StartedAt   time.Time `json:"StartedAt"`
	CompletedAt time.Time `json:"CompletedAt"`

	// === RELATIONS ===
	FlowID string `json:"FlowID"`

	Flow Flow `gorm:"foreignKey:FlowID;references:Name;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	Logs []Log `gorm:"foreignKey:SessionID;"`
}
