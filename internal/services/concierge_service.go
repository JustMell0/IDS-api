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

func (s *ConciergeService) GetInHotelServices() ([]models.InHotelService, error) {
	ctx := context.Background()
	rows, err := s.db.QueryContext(ctx, "SELECT service_id, price, s_name, daily FROM Service NATURAL JOIN InHotelService")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []models.InHotelService
	for rows.Next() {
		var service models.InHotelService
		if err := rows.Scan(&service.ID, &service.Price, &service.Name, &service.Daily); err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (s *ConciergeService) GetTourServices() ([]models.TourService, error) {
	ctx := context.Background()
	rows, err := s.db.QueryContext(ctx, "SELECT service_id, price, s_name, duration, location, transfer FROM Service NATURAL JOIN TourService")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []models.TourService
	for rows.Next() {
		var service models.TourService
		if err := rows.Scan(&service.ID, &service.Price, &service.Name, &service.Duration, &service.Location, &service.Transfer); err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}
