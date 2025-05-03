package services

import (
	"context"
	"database/sql"
)

type LoginService struct {
	db *sql.DB
}

func NewLoginService(db *sql.DB) *LoginService {
	return &LoginService{db: db}
}

func (s *LoginService) Login(id int) (int, error) {
	ctx := context.Background()
	row := s.db.QueryRowContext(ctx, "SELECT guest_id FROM Guest WHERE phone_num = :1", id)
	var guestID int
	if err := row.Scan(&guestID); err != nil {
		if err == sql.ErrNoRows {
			return 0, err // No guest found
		}
		return 0, err
	}
	return guestID, nil
}
