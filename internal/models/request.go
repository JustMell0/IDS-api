package models

type Request struct {
	RequestID    int
	GuestName    string
	GuestSurname string
	ServiceName  string
}

type UserRequest struct {
	RequestID   int    `json:"request_id"`
	ServiceID   int    `json:"service_id"`
	ServiceName string `json:"s_name"`
	Price       int    `json:"price"`
}
