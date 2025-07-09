package main

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

var (
	healthBody  = []byte(`{"status":"ok"}`)
	successBody = []byte(`{"status":"success"}`)
)

func HealthHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.SetBody(healthBody)
}

func PaymentsHandler(ctx *fasthttp.RequestCtx) {
	if !ctx.IsPost() {
		ctx.Error("Method not allowed", fasthttp.StatusMethodNotAllowed)
		return
	}
	var req PaymentRequest
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		ctx.Error("Bad request: "+err.Error(), fasthttp.StatusBadRequest)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.SetBody(successBody)
}
