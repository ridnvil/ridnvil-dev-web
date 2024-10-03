package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenConnectionSQLite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./database/nvil.sqlite3"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
