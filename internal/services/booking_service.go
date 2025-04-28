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

func (s *BookingService) CreateBooking(b models.Booking) (int, error) {
	// Start a single transaction for both operations
	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback() // This will be a no-op if tx.Commit() is called

	// Insert guest record
	var guestID int
	_, err = tx.ExecContext(ctx,
		`INSERT INTO Guest (g_name, surname, phone_num)
     VALUES (:1, :2, :3)
     RETURNING guest_id INTO :4`,
		b.Name, b.Surname, b.Phone, sql.Out{Dest: &guestID})
	if err != nil {
		return 0, err
	}
	var roomNum int
	log.Println("Sending:", b.RoomType)
	err = tx.QueryRowContext(ctx, `SELECT get_room_num_by_type(:1) FROM dual`, b.RoomType).Scan(&roomNum)
	log.Println("Room Got:", roomNum)
	// Insert reservation record using the same transaction
	_, err = tx.ExecContext(context.Background(),
		`INSERT INTO Reservation (guest_id, room_num)
         VALUES (:1, :2)`,
		guestID, roomNum)
	if err != nil {
		return 0, err
	}

	// Commit the transaction once both inserts are successful
	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return int(guestID), nil
}
