package services

import (
	"IDS/api/internal/models"
	"context"
	"database/sql"
	"strings"
)

type RoomTypeService struct {
	db *sql.DB
}

func NewRoomTypeService(db *sql.DB) *RoomTypeService {
	return &RoomTypeService{
		db: db,
	}
}

func (s *RoomTypeService) GetRoomTypes() ([]models.RoomType, error) {
	rows, err := s.db.QueryContext(context.Background(), "SELECT * FROM RoomType")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roomTypes []models.RoomType
	for rows.Next() {
		var room models.RoomType
		if err := rows.Scan(&room.Type, &room.MaxCapacity, &room.Price); err != nil {
			return nil, err
		}
		room.Type = strings.TrimSpace(room.Type)
		roomTypes = append(roomTypes, room)
	}

	return roomTypes, nil
}
