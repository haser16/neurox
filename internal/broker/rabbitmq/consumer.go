package broker_redis

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	channel *amqp091.Channel
}

type HandlerFunc func(ctx context.Context, body []byte) error

func NewConsumer(conn *Connection) (*Consumer, error) {
	ch, err := conn.Conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("create channel: %w", err)
	}

	if err := ch.Qos(10, 0, false); err != nil {
		ch.Close()
		return nil, fmt.Errorf("set QoS: %w", err)
	}
	return &Consumer{channel: ch}, nil
}

func (c *Consumer) Listen(ctx context.Context, queueName, consumerTag string, handler HandlerFunc) error {
	msgs, err := c.channel.ConsumeWithContext(
		ctx,
		queueName,
		consumerTag,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to start consuming: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case msg, ok := <-msgs:
			if !ok {
				return nil
			}

			if err := handler(ctx, msg.Body); err != nil {
				_ = msg.Nack(false, true)
			} else {
				_ = msg.Ack(false)
			}
		}
	}
}

func (c *Consumer) Close() error {
	if c.channel != nil {
		return c.channel.Close()
	}
	return nil
}
