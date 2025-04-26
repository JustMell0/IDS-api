package services

import (
	"IDS/api/internal/models"
	"context"
	"database/sql"
)

type RoomService struct {
	db *sql.DB
}

func NewRoomService(db *sql.DB) *RoomService {
	return &RoomService{db: db}
}

func (s *RoomService) GetRooms() ([]models.Room, error) {
	rows, err := s.db.QueryContext(context.Background(), "SELECT * FROM Room")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room
	for rows.Next() {
		var room models.Room
		if err := rows.Scan(&room.Num, &room.Type); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}
