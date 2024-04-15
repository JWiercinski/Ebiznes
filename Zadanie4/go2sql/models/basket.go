package models

import "gorm.io/gorm"

type Basket struct {
	gorm.Model
	ID    uint64 `gorm:"primaryKey;autoIncrement"`
	USER  string
	COUNT int
	VALUE float64
}

//Skoro w zadaniu nie wymagano relacji z produktem dla koszyka, efekt będzie dość podstawowy
