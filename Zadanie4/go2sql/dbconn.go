package main

import (
	"go2sql/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func dbconn() *gorm.DB {
	mydb, err := gorm.Open(sqlite.Open("godb.db"))
	if err != nil {
		panic("Połączenie z bazą danych nie zostało ustanowione")
	}
	mydb.AutoMigrate(&models.Product{})
	mydb.AutoMigrate(&models.Basket{})
	return mydb
}
