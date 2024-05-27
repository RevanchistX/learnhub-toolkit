package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"learnhub-toolkit/api/database/models"
)

func Db() *gorm.DB {
	dsn := "postgres://default:iy1zXMhxP7kb@ep-dawn-hill-a4qfxq4l-pooler.us-east-1.aws.neon.tech:5432/verceldb"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Migrate() {
	Db().AutoMigrate(&models.Room{})
}
