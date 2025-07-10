package domain

type RouteSummary struct {
	TotalRequests int64   `json:"totalRequests"`
	TotalAmount   float64 `json:"totalAmount"`
}

type PaymentsSummary struct {
	Default  RouteSummary `json:"default"`
	Fallback RouteSummary `json:"fallback"`
}
