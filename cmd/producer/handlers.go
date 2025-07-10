package main

import (
	"encoding/json"

	"github.com/MatheusAbdias/gopay/internal/payment/domain"
	"github.com/MatheusAbdias/gopay/internal/payment/infra/database"
	"github.com/shopspring/decimal"
	amqp "github.com/streadway/amqp"
	"github.com/valyala/fasthttp"
)

func PaymentsSummaryHandler(ctx *fasthttp.RequestCtx, repo *database.PaymentRepository) {
	summary, err := repo.GetPaymentsSummary()
	if err != nil {
		ctx.Error("Failed to fetch summary: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	respBytes, err := json.Marshal(summary)
	if err != nil {
		ctx.Error("Failed to marshal summary: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.SetBody(respBytes)
}

func PaymentsHandler(ctx *fasthttp.RequestCtx, ch *amqp.Channel, repo *database.PaymentRepository) {
	if !ctx.IsPost() {
		ctx.Error("Method not allowed", fasthttp.StatusMethodNotAllowed)
		return
	}

	var paymentRequest PaymentRequest
	if err := json.Unmarshal(ctx.PostBody(), &paymentRequest); err != nil {
		ctx.Error("Invalid request body: "+err.Error(), fasthttp.StatusBadRequest)
		return
	}

	payment, err := domain.NewPayment(paymentRequest.ID, decimal.NewFromFloat(paymentRequest.Amount))
	if err != nil {
		ctx.Error("Invalid payment request: "+err.Error(), fasthttp.StatusBadRequest)
		return
	}

	if err := repo.Save(payment); err != nil {
		ctx.Error("Failed to save payment: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	if err := Publish(ch, *payment); err != nil {
		ctx.Error("Failed to enqueue payment: "+err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusAccepted)
	ctx.SetContentType("application/json")
}
