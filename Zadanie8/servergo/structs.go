package main

import (
	"gorm.io/gorm"
)

type loginJ struct {
	USER     string `json:"USER"`
	PASSWORD string `json:"PASSWORD"`
}

type loginDB struct {
	gorm.Model
	USER     string `gorm:"primaryKey"`
	NAME     string
	EMAIL    string
	PASSWORD string
	TOKEN    string
}

type registerJ struct {
	USER     string `json:"USER"`
	NAME     string `json:"NAME"`
	EMAIL    string `json:"EMAIL"`
	PASSWORD string `json:"PASSWORD"`
}

type responder struct {
	USER  string `json:"USER"`
	TOKEN string `json:"TOKEN"`
}
