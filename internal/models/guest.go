package models

type Guest struct {
	ID       int    `json:"guest_id"`
	Surname  string `json:"surname"`
	Name     string `json:"g_name"`
	PhoneNum int    `json:"phone_num"`
}
