package broker_redis

import (
	"context"
	"fmt"
	core_logger "neurox/internal/core/logger"

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
	log := core_logger.FromContext(ctx)
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

	log.Debug("[WORKER] Started listening on queue")

	for {
		select {
		case <-ctx.Done():
			log.Debug("[WORKER] Stopping listener: context canceled")
			return nil

		case msg, ok := <-msgs:
			if !ok {
				log.Debug("[WORKER] Consumer channel closed by broker")
				return nil
			}

			if err := handler(ctx, msg.Body); err != nil {
				log.Debug("[WORKER] Error handling message. Requeuing...")
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
