package messaging

type Producer interface {
    Publish(exchange, routingKey string, body []byte) error
}

type Consumer interface {
    Consume(queueName string) (<-chan []byte, error)
}
