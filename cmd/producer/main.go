package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MatheusAbdias/gopay/internal/payment/infra/database"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	amqp "github.com/streadway/amqp"
	"github.com/valyala/fasthttp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var db *pgxpool.Pool
var repo *database.PaymentRepository

func init() {
	var err error
	err = godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loadingenv file")
	}
	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err = conn.Channel()
	if err != nil {
		panic(err)
	}
	db_name := os.Getenv("POSTGRES_DB")
	db_user := os.Getenv("POSTGRES_USER")
	db_password := os.Getenv("POSTGRES_PASSWORD")
	db_host := os.Getenv("POSTGRES_HOST")
	db_port := os.Getenv("POSTGRES_PORT")
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", db_user, db_password, db_host, db_port, db_name)
	db, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	repo = database.NewPaymentRepository(db)
}

func main() {
	defer conn.Close()
	defer ch.Close()
	defer db.Close()

	if err := fasthttp.ListenAndServe(":8080", requestHandler); err != nil {
		log.Fatalf("error serving: %v", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/payments":
		PaymentsHandler(ctx, ch, repo)
	case "/payments-summary":
		PaymentsSummaryHandler(ctx, repo)
	default:
		ctx.Error("Not found", fasthttp.StatusNotFound)
	}
}
