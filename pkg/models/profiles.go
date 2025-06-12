package models

import "time"

type Profile struct {
	ID       int       `gorm:"type:int(2);primaryKey;autoIncrement;not null;"`
	Name     string    `gorm:"type:varchar(30);not null;"`
	Position string    `gorm:"type:varchar(100);not null;"`
	Birtdate time.Time `gorm:"type:date;"`
	Bio      string    `gorm:"type:text;"`
	Address  string    `gorm:"type:text;"`
	Phone    string    `gorm:"type:varchar(12);not null;"`
	Email    string    `gorm:"type:varchar(50);not null;"`

	SocialNetworkRef *SocialNetwork `gorm:"foreignKey:ProfileID"`
	ExperincesRef    *Experiences   `gorm:"foreignKey:ProfileID"`
	EducationRef     *Educations    `gorm:"foreignKey:ProfileID"`
	Skill            *Skill         `gorm:"foreignKey:ProfileID"`
}

type SocialNetwork struct {
	ID        int    `gorm:"type:int(3);primaryKey;autoIncrement;not null;"`
	ProfileID int    `gorm:"type:int(2);not null;"`
	Name      string `gorm:"type:varchar(30);not null;"`
	Link      string `gorm:"type:varchar(100);not null;"`
}

type ProfileHome struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Position string    `json:"position"`
	Birtdate time.Time `json:"-"`
	Bio      string    `json:"bio"`
	Address  string    `json:"address"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`

	SocialNetworkRef []SocialNetwork `json:"social_network_ref"`
	ExperincesRef    []Experiences   `json:"experinces_ref"`
	EducationRef     []Educations    `json:"education_ref"`
	SkillRef         []Skill         `json:"skill_ref"`
}
