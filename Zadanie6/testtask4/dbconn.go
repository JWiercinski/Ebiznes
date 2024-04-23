package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testtask4/models"
)

func dbconn() *gorm.DB {
	mydb, err := gorm.Open(sqlite.Open("godb.db"))
	if err != nil {
		panic("Połączenie z bazą danych nie zostało ustanowione")
	}
	mydb.AutoMigrate(&models.Category{})
	otherCategory := new(models.Category)
	otherCategory.ID = 1
	otherCategory.NAME = "INNE"
	if mydb.Model(&models.Category{}).First(&otherCategory).Error != nil {
		mydb.Create(&otherCategory)
	}
	mydb.AutoMigrate(&models.Product{})
	mydb.AutoMigrate(&models.Basket{})
	return mydb
}
