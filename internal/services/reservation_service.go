package services

import (
	"IDS/api/internal/models"
	"context"
	"database/sql"
	"strings"
)

type ReservationService struct {
	db *sql.DB
}

func NewReservationService(db *sql.DB) *ReservationService {
	return &ReservationService{db: db}
}

func (s *ReservationService) GetReservations() ([]models.Reservation, error) {
	ctx := context.Background()
	rows, err := s.db.QueryContext(ctx, "SELECT * FROM Reservation")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservation
	for rows.Next() {
		var reservation models.Reservation
		if err := rows.Scan(&reservation.ID, &reservation.GuestID,
			&reservation.RoomNum, &reservation.CheckInDate,
			&reservation.CheckOutDate, &reservation.TotalPrice,
			&reservation.Status); err != nil {
			return nil, err
		}
		reservation.Status = strings.TrimSpace(reservation.Status)
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}
