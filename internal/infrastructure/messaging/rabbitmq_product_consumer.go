package messaging

import (
	"boilerplate-go/internal/domain/model"
	"boilerplate-go/internal/domain/repository"
	"context"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQProductConsumer struct {
    conn *amqp.Connection
    ch   *amqp.Channel
    repo repository.ProductRepository
}

func NewRabbitMQProductConsumer(amqpURI string, repo repository.ProductRepository) (*RabbitMQProductConsumer, error) {
    conn, err := amqp.Dial(amqpURI)
    if err != nil {
        return nil, err
    }

    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }

    return &RabbitMQProductConsumer{conn: conn, ch: ch, repo: repo}, nil
}

func (c *RabbitMQProductConsumer) Consume(queueName string) (<-chan []byte, error) {
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

func (c *RabbitMQProductConsumer) StartConsumer(queueName string) error {
    msgs, err := c.Consume(queueName)
    if err != nil {
        return err
    }

    go func() {
        for body := range msgs {
            var product model.Product
            if err := json.Unmarshal(body, &product); err != nil {
                log.Printf("Error decoding JSON: %s", err)
                continue
            }

            if err := c.repo.CreateProduct(context.Background(), &product); err != nil {
                log.Printf("Error saving product: %s", err)
            } else {
                log.Printf("Product saved: %v", product)
            }
        }
    }()
    return nil
}
