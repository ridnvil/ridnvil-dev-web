package models

type Educations struct {
	ID              int    `gorm:"type:int(2);primaryKey;autoIncrement;not null;"`
	ProfileID       int    `gorm:"type:int(2);not null;"`
	Institution     string `gorm:"type:varchar(100);not null;"`
	Major           string `gorm:"type:varchar(100);not null;"`
	Description     string `gorm:"type:text;not null;"`
	EducationLength string `gorm:"type:varchar(20);not null;"`
}
