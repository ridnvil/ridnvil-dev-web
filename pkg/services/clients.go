package services

import (
	"gorm.io/gorm"
	"log"
	"ridnvil-dev/pkg/models"
	"ridnvil-dev/pkg/repository"
)

func CreateClient(data models.IPInfo, db *gorm.DB) error {
	if err := repository.CreateCleint(data, db); err != nil {
		log.Println(err)
	}
	return nil
}
