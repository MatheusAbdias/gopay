package domain

type PaymentRepositoryInterface interface {
	Save(payment *Payment) error
	SetProcessedAt(payment *Payment) error
	GetPaymentsSummary() (PaymentsSummary, error)
}
