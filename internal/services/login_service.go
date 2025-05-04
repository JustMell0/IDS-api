package services

import (
	"IDS/api/internal/models"
	"context"
	"database/sql"
)

type LoginService struct {
	db *sql.DB
}

func NewLoginService(db *sql.DB) *LoginService {
	return &LoginService{db: db}
}

func (s *LoginService) Login(id int) (models.Guest, error) {
	ctx := context.Background()
	row := s.db.QueryRowContext(ctx, "SELECT guest_id, surname, g_name, phone_num FROM Guest WHERE phone_num = :1", id)
	var guest models.Guest
	if err := row.Scan(&guest.ID, &guest.Surname, &guest.Name, &guest.PhoneNum); err != nil {
		return guest, err
	}
	return guest, nil
}
