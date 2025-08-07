package models

import "time"

type Flow struct {
	Name      string `gorm:"primaryKey"`
	CreatedAt time.Time

	Sessions []Session `gorm:"foreignKey:FlowId"`
	Logs     []Log     `gorm:"foreignKey:FlowId"`
}
