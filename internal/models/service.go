package models

type InHotelService struct {
	ID    int    `json:"service_id"`
	Price int    `json:"price"`
	Name  string `json:"s_name"`
	Daily string `json:"daily"`
}

type TourService struct {
	ID       int    `json:"service_id"`
	Price    int    `json:"price"`
	Name     string `json:"s_name"`
	Location string `json:"location"`
	Transfer string `json:"transfer"`
	Duration int    `json:"duration"`
}
