package main

import (
	"github.com/valyala/fasthttp"
)

func main() {
	if err := fasthttp.ListenAndServe(":8080", func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/payments":
			PaymentsHandler(ctx)
		case "/health":
			HealthHandler(ctx)
		default:
			ctx.Error("Not found", fasthttp.StatusNotFound)
		}
	}); err != nil {
		panic(err)
	}
}
