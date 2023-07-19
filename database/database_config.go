package database

import (
	"github.com/RafatMeraz/h20/auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type database struct {
	database *gorm.DB
}

func (database *database) Connect() {
	dsn := "root:my_secret_pw@tcp(127.0.0.1:3306)/h2o?charset=utf8mb4&parseTime=True&loc=Local"
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
	err := database.database.AutoMigrate(models.User{})
	if err != nil {
		panic(err)
	}
}

var Database database
