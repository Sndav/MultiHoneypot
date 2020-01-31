package models

import "time"

const (
	TableVictim = "victim"
)

type Victim struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	FromIP    string    `gorm:"size:50" json:"from_ip"`
	Service   string    `gorm:"size:50" json:"service"`
	Payload   string    `gorm:"type:longtext" json:"payload"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Victim) TableName() string {
	return TableVictim
}
