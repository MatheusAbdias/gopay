package main

import (
	"encoding/json"

	"github.com/MatheusAbdias/gopay/internal/payment/domain"
	amqp "github.com/streadway/amqp"
)

func Publish(ch *amqp.Channel, payment domain.Payment) error {
	body, err := json.Marshal(payment)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
