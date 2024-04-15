package models

type ProductJ struct {
	NAME       string  `json:"NAME"`
	PRICE      float64 `json:"PRICE"`
	CATEGORYID uint64  `json:"CATEGORYID"`
}
