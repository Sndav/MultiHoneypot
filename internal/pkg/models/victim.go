package models

import "time"

const (
	TableVictim = "victim"
)

type Victim struct {
	ID        uint   `gorm:"primary_key"`
	FromIP    string `gorm:"size:50"`
	Service   string `gorm:"size:50"`
	Payload   string `gorm:"type:longtext"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Victim) TableName() string {
	return TableVictim
}