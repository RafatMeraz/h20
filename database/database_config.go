package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	database *gorm.DB
}

func (database *Database) Connect() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database.database = db
	database.migrations()
}

func (database *Database) Instance() *gorm.DB {
	return database.database
}

func (database *Database) migrations() {
	database.database.AutoMigrate()
}
