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

func (s *ReservationService) EditReservation(guestID int, checkInDate string, checkOutDate string) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "UPDATE Reservation SET check_in_date = :1, check_out_date = :2 WHERE guest_id = :3", checkInDate, checkOutDate, guestID)
	return err
}

func (s *ReservationService) GetUserReservations(id int) ([]models.UserReservation, error) {
	ctx := context.Background()
	rows, err := s.db.QueryContext(ctx, "SELECT reservation_id, guest_id, room_num, room_type, check_in_date, check_out_date, total_price, status FROM Reservation NATURAL JOIN Room WHERE guest_id = :1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.UserReservation
	for rows.Next() {
		var reservation models.UserReservation
		if err := rows.Scan(&reservation.ID, &reservation.GuestID,
			&reservation.Room.RoomNum, &reservation.Room.Type, &reservation.CheckInDate,
			&reservation.CheckOutDate, &reservation.TotalPrice,
			&reservation.Status); err != nil {
			return nil, err
		}
		reservation.Status = strings.TrimSpace(reservation.Status)
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}
