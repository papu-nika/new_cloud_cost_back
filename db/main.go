package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "host=localhost user=postgres password=adminadmin dbname=cloud_cost port=5432 sslmode=disable TimeZone=Asia/Tokyo"

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
