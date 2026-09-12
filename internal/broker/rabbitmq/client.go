package broker_redis

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type Connection struct {
	Conn *amqp091.Connection
}

func NewConnection(config Config) (*Connection, error) {
	connectionString := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		config.User,
		config.Password,
		config.Host,
		config.Port,
	)
	conn, err := amqp091.Dial(connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to rabbitmq: %w", err)
	}

	return &Connection{
		Conn: conn,
	}, nil
}

func (c *Connection) Close() error {
	if c.Conn != nil && !c.Conn.IsClosed() {
		return c.Conn.Close()
	}
	return nil
}
