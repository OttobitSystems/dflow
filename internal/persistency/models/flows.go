// Package models: are all the models for repository
package models

import "time"

type Flow struct {
	Name      string    `json:"Name" gorm:"primaryKey"`
	CreatedAt time.Time `json:"CreatedAt"`

	Sessions []Session `gorm:"foreignKey:FlowID"`
	Logs     []Log     `gorm:"foreignKey:FlowID"`
}
