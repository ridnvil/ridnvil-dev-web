package models

import "time"

type Profile struct {
	ID       int       `gorm:"type:int(2);primaryKey;autoIncrement;not null;"`
	Name     string    `gorm:"type:varchar(30);not null;"`
	Birtdate time.Time `gorm:"type:date;"`
	Bio      string    `gorm:"type:text;"`
	Address  string    `gorm:"type:text;"`
	Phone    string    `gorm:"type:varchar(12);not null;"`
	Email    string    `gorm:"type:varchar(50);not null;"`

	SocialNetworkRef *SocialNetwork `gorm:"foreignKey:ProfileID"`
	ExperincesRef    *Experinces    `gorm:"foreignKey:ProfileID"`
	EducationRef     *Educations    `gorm:"foreignKey:ProfileID"`
}

type SocialNetwork struct {
	ID        int    `gorm:"type:int(3);primaryKey;autoIncrement;not null;"`
	ProfileID int    `gorm:"type:int(2);not null;"`
	Link      string `gorm:"type:varchar(100);not null;"`
}
