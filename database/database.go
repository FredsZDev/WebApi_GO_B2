package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dtb *gorm.DB

func StartDB() {
	str := "host=localhost port=25000 user=admin dbname=games sslmode=disable password=123456"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("error:", err)

	}

	dtb = database

	config, _ := dtb.DB()

	config.SetConnMaxIdleTime(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB {

	return dtb
}
