package models

type Experiences struct {
	ID          int    `gorm:"type:int(11);primaryKey;autoIncrement;not null;" json:"id"`
	ProfileID   int    `gorm:"type:int(2);not null;" json:"profile_id"`
	Position    string `gorm:"type:varchar(100);not null;" json:"position"`
	CompanyName string `gorm:"type:varchar(100);not null;" json:"company_name"`
	Description string `gorm:"type:text;" json:"description"`
	StartYears  string `gorm:"type:varchar(4);not null;default:'2018'" json:"start_years"`
	UntilYears  string `gorm:"type:varchar(4);not null;default:'Now'" json:"until_years"`
}
