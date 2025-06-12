package models

type Skill struct {
	ID          int    `gorm:"type:int(11);primaryKey;autoIncrement;not null;" json:"id"`
	ProfileID   int    `gorm:"type:int(2);not null;" json:"profile_id"` // Foreign key to Profile
	Name        string `gorm:"type:varchar(100);not null;" json:"name"`
	Description string `gorm:"type:text;" json:"description"`
	Level       string `gorm:"type:varchar(50);not null;" json:"level"` // e.g., Beginner, Intermediate, Advanced
}
