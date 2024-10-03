package services

import (
	"log"
	"ridnvil-dev/pkg/models"
	"ridnvil-dev/pkg/repository"
)

func CreateClient(data models.IPInfo) error {
	if err := repository.CreateCleint(data); err != nil {
		log.Println(err)
	}
	return nil
}
