package main

type PaymentRequest struct {
	ID     string  `json:"correlationId"`
	Amount float64 `json:"amount"`
}
