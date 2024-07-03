package messaging

import (
	"github.com/streadway/amqp"
)

type RabbitMQProducer struct {
    conn *amqp.Connection
    ch   *amqp.Channel
}

func NewRabbitMQProducer(amqpURI string) (*RabbitMQProducer, error) {
    conn, err := amqp.Dial(amqpURI)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    return &RabbitMQProducer{conn: conn, ch: ch}, nil
}

func (p *RabbitMQProducer) Publish(exchange, routingKey string, body []byte) error {
    return p.ch.Publish(
        exchange,   // exchange
        routingKey, // routing key
        false,      // mandatory
        false,      // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        })
}
