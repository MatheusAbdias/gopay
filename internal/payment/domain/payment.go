package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Payment struct {
	ID          string          `db:"id"`
	Amount      decimal.Decimal `db:"amount"`
	Processor   string          `db:"processor"`
	CreatedAt   time.Time       `db:"created_at"`
	ProcessedAt *time.Time      `db:"processed_at"`
}

func NewPayment(id string, amount decimal.Decimal) (*Payment, error) {
	payment := &Payment{
		ID:        id,
		Amount:    amount,
		CreatedAt: time.Now(),
	}

	err := payment.IsValid()
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (p *Payment) IsValid() error {
	if p.Amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("invalid amount")
	}

	if _, err := uuid.Parse(p.ID); err != nil {
		return errors.New("invalid uuid")
	}

	return nil
}
