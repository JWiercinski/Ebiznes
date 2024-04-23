package models

type CategoryJ2 struct {
	ID       int         `json:"ID"`
	NAME     string      `json:"NAME"`
	PRODUCTS []ProductJ2 `json:"PRODUCTS"`
}
