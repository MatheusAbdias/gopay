package domain

import (
	"testing"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestGivenNegativeAmount_WhenCreateANewPayment_ThenShouldReceiveAnError(t *testing.T) {
	payment := Payment{ID: uuid.New().String(), Amount: decimal.NewFromFloat(-10.0)}
	assert.Error(t, payment.IsValid(), "invalid amount")
}

func TestGivenInvalidUUID_WhenCreateANewPayment_ThenShouldReceiveAnError(t *testing.T) {
	payment := Payment{ID: "invalid", Amount: decimal.NewFromFloat(10.0)}
	assert.Error(t, payment.IsValid(), "invalid uuid")
}
