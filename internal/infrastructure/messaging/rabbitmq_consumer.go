package messaging

import (
	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
    conn *amqp.Connection
    ch   *amqp.Channel
}

func NewRabbitMQConsumer(amqpURI, queueName string) (*RabbitMQConsumer, error) {
    conn, err := amqp.Dial(amqpURI)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    _, err = ch.QueueDeclare(
        queueName, // name
        true,      // durable
        false,     // delete when unused
        false,     // exclusive
        false,     // no-wait
        nil,       // arguments
    )
    if err != nil {
        return nil, err
    }

    return &RabbitMQConsumer{conn: conn, ch: ch}, nil
}

func (c *RabbitMQConsumer) Consume(queueName string) (<-chan []byte, error) {
    msgs, err := c.ch.Consume(
        queueName, // queue
        "",        // consumer
        true,      // auto-ack
        false,     // exclusive
        false,     // no-local
        false,     // no-wait
        nil,       // args
    )
    if err != nil {
        return nil, err
    }

    out := make(chan []byte)
    go func() {
        for d := range msgs {
            out <- d.Body
        }
        close(out)
    }()
    return out, nil
}
