package database

import (
	"context"

	"github.com/MatheusAbdias/gopay/internal/payment/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository struct {
	Pool *pgxpool.Pool
}

func NewPaymentRepository(pool *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{Pool: pool}
}

func (r *PaymentRepository) Save(payment *domain.Payment) error {
	_, err := r.Pool.Exec(context.Background(),
		"INSERT INTO payments (id, amount, processor, created_at, processed_at) VALUES ($1, $2, $3, $4, $5)",
		payment.ID, payment.Amount, payment.Processor, payment.CreatedAt, payment.ProcessedAt,
	)

	return err
}

func (r *PaymentRepository) SetProcessedAt(payment *domain.Payment) error {
	_, err := r.Pool.Exec(context.Background(),
		"UPDATE payments SET processed_at = $1 WHERE id = $2",
		payment.ProcessedAt, payment.ID,
	)

	return err
}

func (r *PaymentRepository) GetPaymentsSummary() (domain.PaymentsSummary, error) {
	query := `
	SELECT
	  SUM(CASE WHEN processor = 'default' THEN 1 ELSE 0 END) AS default_requests,
	  SUM(CASE WHEN processor = 'default' THEN amount ELSE 0 END) AS default_amount,
	  SUM(CASE WHEN processor = 'fallback' THEN 1 ELSE 0 END) AS fallback_requests,
	  SUM(CASE WHEN processor = 'fallback' THEN amount ELSE 0 END) AS fallback_amount
	FROM payments;
	`
	var defaultReq, fallbackReq int64
	var defaultAmt, fallbackAmt float64
	err := r.Pool.QueryRow(context.Background(), query).Scan(&defaultReq, &defaultAmt, &fallbackReq, &fallbackAmt)
	if err != nil {
		return domain.PaymentsSummary{}, err
	}
	return domain.PaymentsSummary{
		Default: domain.RouteSummary{
			TotalRequests: defaultReq,
			TotalAmount:   defaultAmt,
		},
		Fallback: domain.RouteSummary{
			TotalRequests: fallbackReq,
			TotalAmount:   fallbackAmt,
		},
	}, nil
}
