// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package logger_di

import (
	"context"
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/di/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/internal/services/logger/application"
	"github.com/shortlink-org/shortlink/internal/services/logger/infrastructure/mq"
	"go.opentelemetry.io/otel/trace"
)

// Injectors from wire.go:

func InitializeLoggerService() (*LoggerService, func(), error) {
	context, cleanup, err := ctx.New()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := logger_di.New(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	configConfig, err := config.New()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	monitoringMonitoring, cleanup3, err := monitoring.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	tracerProvider, cleanup4, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint, err := profiling.New(context, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup5, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewLoggerApplication(logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	mq, err := mq_di.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	event, err := InitLoggerMQ(context, logger, mq, service)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	loggerService, err := NewLoggerService(logger, configConfig, monitoringMonitoring, tracerProvider, pprofEndpoint, autoMaxProAutoMaxPro, service, event)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return loggerService, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type LoggerService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Delivery
	loggerMQ *logger_mq.Event

	// Application
	loggerService *logger_application.Service
}

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(di.DefaultSet, mq_di.New, InitLoggerMQ,

	NewLoggerApplication,

	NewLoggerService,
)

func InitLoggerMQ(ctx2 context.Context, log logger.Logger, dataBus mq.MQ, service *logger_application.Service) (*logger_mq.Event, error) {
	loggerMQ, err := logger_mq.New(dataBus, log, service)
	if err != nil {
		return nil, err
	}

	return loggerMQ, nil
}

func NewLoggerApplication(log logger.Logger) (*logger_application.Service, error) {
	loggerService, err := logger_application.New(log)
	if err != nil {
		return nil, err
	}

	return loggerService, nil
}

func NewLoggerService(

	log logger.Logger, config2 *config.Config, monitoring2 *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,

	loggerService *logger_application.Service,

	loggerMQ *logger_mq.Event,
) (*LoggerService, error) {
	return &LoggerService{

		Log:    log,
		Config: config2,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		loggerService: loggerService,

		loggerMQ: loggerMQ,
	}, nil
}
