package models

type Game struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Price    string `json:"price"`
	IsOnline bool   `json:"isonline"`
}
