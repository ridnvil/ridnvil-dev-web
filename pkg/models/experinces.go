package models

type Experinces struct {
	ID          int    `gorm:"type:int(11);primaryKey;autoIncrement;not null;"`
	ProfileID   int    `gorm:"type:int(2);not null;"`
	Position    string `gorm:"type:varchar(100);not null;"`
	Description string `gorm:"type:text;"`
	StartYears  string `gorm:"type:varchar(4);not null;default:'2018'"`
	UntilYears  string `gorm:"type:varchar(4);not null;default:'Now'"`
}
