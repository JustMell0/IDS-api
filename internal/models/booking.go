package models

type Booking struct {
	Name         string `json:"g_name"`
	Surname      string `json:"surname"`
	Phone        int    `json:"phone_num"`
	RoomType     string `json:"room_type"`
	CheckInDate  string `json:"check_in_date"`
	CheckOutDate string `json:"check_out_date"`
}
