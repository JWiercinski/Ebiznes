package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID    uint64 `gorm:"primaryKey;autoIncrement"`
	NAME  string
	PRICE float64
}
