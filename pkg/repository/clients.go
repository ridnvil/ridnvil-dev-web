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

	var ipstring string

	if errcheck := db.Model(models.IPInfo{}).Select("ip").Where("ip = ?", data.IP).Scan(&ipstring).Error; errcheck != nil {
		return errcheck
	}

	if ipstring == "" {
		if err := db.Create(&data).Error; err != nil {
			return err
		}
	}

	if errupdate := db.Where("ip = ?", data.IP).Save(&data).Error; errupdate != nil {
		return errupdate
	}

	return nil
}

func GetListClient() ([]models.IPInfo, error) {
	db, err := database.OpenConnectionSQLite()
	if err != nil {
		return []models.IPInfo{}, err
	}

	var list []models.IPInfo
	if errget := db.Find(&list).Error; errget != nil {
		return []models.IPInfo{}, errget
	}

	return list, nil
}
