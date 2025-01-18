package database

import (
	"github.com/caarlos0/env/v11"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"ridnvil-dev/pkg/environtments"
)

func OpenConnectionSQLite() (*gorm.DB, error) {
	var envir environtments.Env

	if err := env.Parse(&envir); err != nil {
		log.Println(err)
	}

	db, err := gorm.Open(mysql.Open(envir.Host), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoCreateDatabase() error {
	var envir environtments.Env
	if err := env.Parse(&envir); err != nil {
		log.Println(err)
	}

	db, err := gorm.Open(mysql.Open(envir.HostOnly), &gorm.Config{})
	if err != nil {
		return err
	}

	dbname := "ridnvil"
	var exists []string

	query := "SHOW DATABASES"
	err = db.Raw(query).Scan(&exists).Error
	if err != nil {
		return err
	}

	if !checkDb(exists) {
		if errcreate := db.Exec("CREATE DATABASE " + dbname).Error; errcreate != nil {
			return errcreate
		}
		log.Println("Database created")
	} else {
		log.Println("Database found")
	}
	return nil
}

func checkDb(list []string) bool {
	for _, dbname := range list {
		if "ridnvil" == dbname {
			return true
		}
	}
	return false
}
