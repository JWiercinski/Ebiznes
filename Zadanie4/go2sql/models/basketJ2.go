package models

type BasketJ2 struct {
	ID    int     `json:"ID"`
	USER  string  `json:"USER"`
	COUNT int     `json:"COUNT"`
	VALUE float64 `json:"VALUE"`
}
