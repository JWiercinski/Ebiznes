package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	NAME     string `gorm:"unique"`
	PRODUCTS []Product
}
