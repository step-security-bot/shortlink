//go:generate wire
//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

/*
Logger Service DI-package
*/
package logger_di

import (
	"context"

	"github.com/google/wire"

	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	mq_di "github.com/shortlink-org/shortlink/internal/di/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq/v1"
	logger_application "github.com/shortlink-org/shortlink/internal/services/logger/application"
	logger_mq "github.com/shortlink-org/shortlink/internal/services/logger/infrastructure/mq"
)

type LoggerService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Delivery
	loggerMQ *logger_mq.Event

	// Application
	loggerService *logger_application.Service
}

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(
	di.DefaultSet,
	mq_di.New,

	// Delivery
	InitLoggerMQ,

	// Applications
	NewLoggerApplication,

	NewLoggerService,
)

func InitLoggerMQ(ctx context.Context, log logger.Logger, mq v1.MQ, service *logger_application.Service) (*logger_mq.Event, error) {
	loggerMQ, err := logger_mq.New(mq, log, service)
	if err != nil {
		return nil, err
	}

	return loggerMQ, nil
}

func NewLoggerApplication(logger logger.Logger) (*logger_application.Service, error) {
	loggerService, err := logger_application.New(logger)
	if err != nil {
		return nil, err
	}

	return loggerService, nil
}

func NewLoggerService(
	// Common
	log logger.Logger,
	config *config.Config,

	// Application
	loggerService *logger_application.Service,

	// Delivery
	loggerMQ *logger_mq.Event,
) (*LoggerService, error) {
	return &LoggerService{
		// Common
		Log:    log,
		Config: config,

		// Application
		loggerService: loggerService,

		// Delivery
		loggerMQ: loggerMQ,
	}, nil
}

func InitializeLoggerService() (*LoggerService, func(), error) {
	panic(wire.Build(LoggerSet))
}
