package models

type UnconfirmedPayment struct {
	ID            int    `json:"payment_id"`
	Guest         Guest  `json:"guest"`
	ReservationID int    `json:"reservation_id"`
	Time          string `json:"time"`
	Amount        int    `json:"amount_paid"`
	Type          string `json:"type"`
}
