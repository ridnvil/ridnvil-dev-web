package repository

import (
	"ridnvil-dev/pkg/database"
	"ridnvil-dev/pkg/models"
)

func CreateCleint(data models.IPInfo) error {
	db, err := database.OpenConnectionSQLite()
	if err != nil {
		return err
	}

	if err := db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}
