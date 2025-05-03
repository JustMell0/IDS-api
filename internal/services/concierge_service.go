package services

import (
	"IDS/api/internal/models"
	"context"
	"database/sql"
)

type ConciergeService struct {
	db *sql.DB
}

func NewConciergeService(db *sql.DB) *ConciergeService {
	return &ConciergeService{db: db}
}

func (s *ConciergeService) GetRequests() ([]models.Request, error) {
	ctx := context.Background()
	// TODO: add WHERE time is start of UNIX epoch
	rows, err := s.db.QueryContext(ctx, "SELECT request_id, g_name, surname, s_name FROM Request NATURAL JOIN Guest NATURAL JOIN Service")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []models.Request
	for rows.Next() {
		var request models.Request
		if err := rows.Scan(&request.RequestID, &request.GuestName,
			&request.GuestSurname, &request.ServiceName); err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}
	return requests, nil

}

func (s *ConciergeService) AcceptRequest(requestID int) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "UPDATE Request SET time = TO_CHAR(SYSDATE, 'YYYY-MM-DD HH24:MI:SS') WHERE request_id = :1", requestID)
	return err
}

func (s *ConciergeService) RejectRequest(requestID int) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "DELETE FROM Request WHERE request_id = :1", requestID)
	return err
}
