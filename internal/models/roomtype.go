package models

type RoomType struct {
	Type        string `json:"type"`
	MaxCapacity int    `json:"max_capacity"`
	Price       int    `json:"price"`
}
