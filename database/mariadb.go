package database

import (
	"github.com/goroutine/template/log"
	"github.com/goroutine/template/utils/environnement"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Maria *gorm.DB

func StartMariaClient() (err error) {
	Maria, err = gorm.Open(mysql.Open(environnement.GetMariaDsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Logger.Fatal("Could not connect to Maria: ", err)
	}
	if err = Maria.AutoMigrate(); err != nil {
		log.Logger.Fatal("Could not migrate tables: ", err)
	}

	db, err := Maria.DB()
	if err != nil {
		log.Logger.Fatal("Could not get Maria db: ", err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(15)
	log.Logger.Info("Maria client started")
	return
}
