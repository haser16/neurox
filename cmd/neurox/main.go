package main

import (
	"context"
	"fmt"
	auth_jwt "neurox/internal/auth/jwt"
	broker_redis "neurox/internal/broker/rabbitmq"
	core_logger "neurox/internal/core/logger"
	core_pgx_pool "neurox/internal/core/repository/postgres/pool/pgx"
	core_middleware "neurox/internal/core/transport/http/middleware"
	core_http_server "neurox/internal/core/transport/http/server"
	users_postgres_repository "neurox/internal/features/users/repository/postgres"
	users_service "neurox/internal/features/users/service"
	users_transport_http "neurox/internal/features/users/transport/http"
	storage_s3 "neurox/internal/storage/s3"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	logger.Debug("application time zone", zap.Any("zone", time.Local))

	logger.Debug("initialize postgres connection pool")
	pool, err := core_pgx_pool.NewPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initialize token service")
	tokenService := auth_jwt.NewAuth(auth_jwt.NewConfigMust())

	logger.Debug("initialize s3 service")
	s3client, err := storage_s3.NewClient(ctx, storage_s3.NewConfigMust())

	if err != nil {
		logger.Fatal(
			"failed to initialize s3 client",
			zap.Error(err),
		)
	}
	s3Storage := storage_s3.NewStorage(s3client, storage_s3.NewConfigMust())

	logger.Debug("initialize rabbitmq service")
	rabbitmqConnection, err := broker_redis.NewConnection(broker_redis.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to initialize rabbitmq connection", zap.Error(err))
	}
	defer rabbitmqConnection.Close()

	publisher, err := broker_redis.NewPublisher(rabbitmqConnection)
	if err != nil {
		logger.Fatal("failed to initialize rabbitmq publisher", zap.Error(err))
	}
	defer publisher.Close()

	if err := publisher.InitQueue(broker_redis.QueueEmailTasks); err != nil {
		logger.Fatal("failed to initialize rabbitmq queue", zap.Error(err))
	}

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository, tokenService, s3Storage, publisher)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("initializing HTTP server")

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_middleware.RequestID(),
		core_middleware.Logger(logger),
		core_middleware.Trace(),
		core_middleware.Panic(),
	)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(
		&core_http_server.APIVersion1,
	)

	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
