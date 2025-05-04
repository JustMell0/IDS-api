package services

import (
	"IDS/api/internal/models"
	"context"
	"database/sql"
)

type PaymentService struct {
	db *sql.DB
}

func NewPaymentService(db *sql.DB) *PaymentService {
	return &PaymentService{db: db}
}

func (s *PaymentService) GetUnconfirmedPayments() ([]models.UnconfirmedPayment, error) {
	ctx := context.Background()
	rows, err := s.db.QueryContext(ctx, `
		SELECT payment_id, guest_id, g_name, surname, phone_num, reservation_id, time, amount_paid, type
		FROM Payment NATURAL JOIN Guest WHERE status = 'UnConfirmed'`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.UnconfirmedPayment
	for rows.Next() {
		var payment models.UnconfirmedPayment
		if err := rows.Scan(&payment.ID, &payment.Guest.ID,
			&payment.Guest.Name, &payment.Guest.Surname, &payment.Guest.PhoneNum,
			&payment.ReservationID, &payment.Time, &payment.Amount,
			&payment.Type); err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (s *PaymentService) ConfirmPayment(id int) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "UPDATE Payment SET status = 'Confirmed' WHERE payment_id = :1", id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PaymentService) RejectPayment(id int) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "DELETE FROM Payment WHERE payment_id = :1", id)
	if err != nil {
		return err
	}
	return nil
}
