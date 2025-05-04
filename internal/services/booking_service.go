package services

import (
	"IDS/api/internal/models"
	"context"
	"database/sql"
	"log"
)

type BookingService struct {
	db *sql.DB
}

func NewBookingService(db *sql.DB) *BookingService {
	return &BookingService{db: db}
}

func (s *BookingService) CreateBooking(b models.Booking) (models.Guest, error) {
	// Start a single transaction for both operations
	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return models.Guest{}, err
	}
	defer tx.Rollback()

	// Insert guest record
	var guestID int
	_, err = tx.ExecContext(ctx,
		`INSERT INTO Guest (g_name, surname, phone_num)
					VALUES (:1, :2, :3)
					RETURNING guest_id INTO :4`,
		b.Name, b.Surname, b.Phone, sql.Out{Dest: &guestID})
	if err != nil {
		return models.Guest{}, err
	}
	var roomNum int
	log.Println("Sending:", b.RoomType)
	err = tx.QueryRowContext(ctx, `SELECT get_room_num_by_type(:1) FROM dual`, b.RoomType).Scan(&roomNum)
	log.Println("Room Got:", roomNum)
	// Insert reservation record using the same transaction
	_, err = tx.ExecContext(context.Background(),
		`INSERT INTO Reservation (guest_id, room_num, check_in_date, check_out_date)
									VALUES (:1, :2, :3, :4)`,
		guestID, roomNum, b.CheckInDate, b.CheckOutDate)
	if err != nil {
		return models.Guest{}, err
	}

	// Commit the transaction once both inserts are successful
	if err = tx.Commit(); err != nil {
		return models.Guest{}, err
	}

	// Create and return the complete guest info
	guest := models.Guest{
		ID:       guestID,
		Name:     b.Name,
		Surname:  b.Surname,
		PhoneNum: b.Phone,
	}

	return guest, nil
}
