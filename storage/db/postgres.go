package db

import (
	"log"

	"github.com/lai0xn/docsoft/config"
	"github.com/lai0xn/docsoft/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate() {
	DB.Debug().AutoMigrate(models.User{})
	DB.AutoMigrate(models.File{})
	DB.AutoMigrate(models.Folder{})
	DB.AutoMigrate(models.Workspace{})
	DB.AutoMigrate(models.Role{})
	DB.AutoMigrate(models.TimeCapsule{})
}

func Connect() {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	log.Println("DB Connected ....")
	Migrate()
}
