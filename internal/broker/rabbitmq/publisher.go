package broker_redis

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	channel *amqp091.Channel
}

func NewPublisher(conn *Connection) (*Publisher, error) {
	ch, err := conn.Conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open publisher channel: %w", err)
	}

	if err := ch.Confirm(false); err != nil {
		ch.Close()
		return nil, fmt.Errorf("failed to put channel into confirm mode: %w", err)
	}

	return &Publisher{
		channel: ch,
	}, nil
}

func (p *Publisher) InitQueue(name string) error {
	_, err := p.channel.QueueDeclare(
		name,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue %s: %w", name, err)
	}
	return nil
}

func (p *Publisher) Publish(ctx context.Context, exchange, routingKey string, body []byte) error {
	if p == nil {
		return fmt.Errorf("publisher is nil")
	}

	if p.channel == nil {
		return fmt.Errorf("publisher channel is nil")
	}

	err := p.channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp091.Persistent,
			Body:         body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}

func (p *Publisher) Close() error {
	if p.channel != nil {
		return p.channel.Close()
	}
	return nil
}
