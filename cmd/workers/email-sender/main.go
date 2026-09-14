package main

import (
	"context"
	"fmt"
	email_sender "neurox/cmd/workers/email-sender/sender"
	broker_redis "neurox/internal/broker/rabbitmq"
	core_logger "neurox/internal/core/logger"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("Failed to initialize logger: ", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initialize rabbitmq service")
	rabbitmqConnection, err := broker_redis.NewConnection(broker_redis.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to initialize rabbitmq connection", zap.Error(err))
	}
	defer rabbitmqConnection.Close()

	consumer, err := broker_redis.NewConsumer(rabbitmqConnection)
	if err != nil {
		logger.Fatal("failed to initialize rabbitmq consumer", zap.Error(err))
	}
	defer consumer.Close()

	logger.Debug("initialize sender service")
	sender, err := email_sender.NewSender(email_sender.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to initialize sender service", zap.Error(err))
	}

	if err := consumer.Listen(
		ctx,
		broker_redis.QueueEmailTasks,
		"consumer",
		sender.SendMessage,
	); err != nil {
		logger.Fatal("consumer stopped", zap.Error(err))
	}
}
