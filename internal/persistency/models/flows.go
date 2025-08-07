// Package models: are all the models for repository
package models

import "time"

type Flow struct {
	Name      string `gorm:"primaryKey"`
	CreatedAt time.Time

	Sessions []Session `gorm:"foreignKey:FlowID"`
	Logs     []Log     `gorm:"foreignKey:FlowID"`
}
