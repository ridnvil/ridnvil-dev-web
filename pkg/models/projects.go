package models

import "time"

type Project struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"size:100;not null" json:"name"`
	Company     string    `gorm:"size:100;not null" json:"company"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Length      string    `gorm:"size:100;not null" json:"length"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}
