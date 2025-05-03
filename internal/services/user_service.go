package services

import (
	"IDS/api/internal/models"
	"context"
	"database/sql"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetUserRequests(id int) ([]models.UserRequest, error) {
	ctx := context.Background()
	// TODO: add WHERE time is start of UNIX epoch
	rows, err := s.db.QueryContext(ctx, "SELECT request_id, service_id, s_name, price FROM Request NATURAL JOIN Guest NATURAL JOIN Service WHERE guest_id = :1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []models.UserRequest
	for rows.Next() {
		var request models.UserRequest
		if err := rows.Scan(&request.RequestID, &request.ServiceID,
			&request.ServiceName, &request.Price); err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}
	return requests, nil
}

func (s *UserService) NewRequest(guestID int, serviceID int) (int, error) {
	ctx := context.Background()
	var requestID int
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO Request (guest_id, service_id)
			VALUES (:1, :2)
			RETURNING request_id INTO :3`, guestID, serviceID, sql.Out{Dest: &requestID})
	if err != nil {
		return 0, err
	}
	return requestID, nil
}
