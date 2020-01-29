package models

import "time"

const (
	TableLog = "logs"
)

type Logs struct {
	ID        uint   `gorm:"primary_key"`
	LogType   string `gorm:"size:50"`
	Log       string `gorm:"type:longtext"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Logs) TableName() string {
	return TableLog
}