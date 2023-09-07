package database

import (
	"fmt"
	"github.com/RafatMeraz/h20/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type database struct {
	database *gorm.DB
}

func (database *database) Connect() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.database = db
	database.migrations()
}

func (database *database) Instance() *gorm.DB {
	return database.database
}

func (database *database) migrations() {
	if err := database.database.AutoMigrate(models.User{}); err != nil {
		panic(err)
	}
	if err := database.database.AutoMigrate(models.WaterTrack{}); err != nil {
		panic(err)
	}
}

var Database database
