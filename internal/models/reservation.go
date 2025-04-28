package models

type Reservation struct {
	ID           int    `json:"reservation_id"`
	GuestID      int    `json:"guest_id"`
	RoomNum      int    `json:"room_num"`
	CheckInDate  string `json:"check_in_date"`
	CheckOutDate string `json:"check_out_date"`
	TotalPrice   int    `json:"total_price"`
	Status       string `json:"status"`
}
