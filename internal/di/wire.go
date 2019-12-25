//go:generate wire
//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"
	"time"

	"github.com/google/wire"
	"github.com/heptiolabs/healthcheck"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/internal/mq/kafka"
	"github.com/batazor/shortlink/internal/store"
	"github.com/batazor/shortlink/internal/traicing"
)

// Service - heplers
type Service struct {
	Log    logger.Logger
	Tracer opentracing.Tracer
	// TracerClose func()
	DB         store.DB
	MQ         mq.MQ
	Monitoring *http.ServeMux
}

// InitStore return store
func InitStore(ctx context.Context, log logger.Logger) (store.DB, error) {
	var st store.Store
	db := st.Use(ctx, log)
	return db, nil
}

func InitLogger(ctx context.Context) (logger.Logger, error) {
	viper.SetDefault("LOG_LEVEL", logger.INFO_LEVEL)
	viper.SetDefault("LOG_TIME_FORMAT", time.RFC3339Nano)

	conf := logger.Configuration{
		Level:      viper.GetInt("LOG_LEVEL"),
		TimeFormat: viper.GetString("LOG_TIME_FORMAT"),
	}

	log, err := logger.NewLogger(logger.Zap, conf)
	if err != nil {
		return nil, err
	}

	return log, nil
}

func InitTracer(ctx context.Context) (opentracing.Tracer, error) {
	viper.SetDefault("TRACER_SERVICE_NAME", "ShortLink")
	viper.SetDefault("TRACER_URI", "localhost:6831")

	config := traicing.Config{
		ServiceName: viper.GetString("TRACER_SERVICE_NAME"),
		URI:         viper.GetString("TRACER_URI"),
	}

	// TODO: add close func to return
	tracer, _, err := traicing.Init(config)
	if err != nil {
		return nil, err
	}

	// Add tracer to context
	ctx = traicing.WithTraicer(ctx, tracer)

	return tracer, nil
}

func InitMonitoring() *http.ServeMux {
	// Create a new Prometheus registry
	registry := prometheus.NewRegistry()

	// Create a metrics-exposing Handler for the Prometheus registry
	// The healthcheck related metrics will be prefixed with the provided namespace
	health := healthcheck.NewMetricsHandler(registry, "common")

	// Our app is not happy if we've got more than 100 goroutines running.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	// Create an "common" listener
	commonMux := http.NewServeMux()

	// Expose prometheus metrics on /metrics
	commonMux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	// Expose a liveness check on /live
	commonMux.HandleFunc("/live", health.LiveEndpoint)

	// Expose a readiness check on /ready
	commonMux.HandleFunc("/ready", health.ReadyEndpoint)

	return commonMux
}

func InitMQ(ctx context.Context) (mq.MQ, error) {
	viper.SetDefault("MQ_ENABLED", "false")

	if viper.GetBool("MQ_ENABLED") {
		var service mq.MQ
		service = &kafka.Kafka{}

		if err := service.Init(ctx); err != nil {
			return nil, err
		}

		return service, nil
	}

	return nil, nil
}

func NewFullService(log logger.Logger, mq mq.MQ, monitoring *http.ServeMux, tracer opentracing.Tracer, db store.DB) (*Service, error) {
	return &Service{
		Log:    log,
		MQ:     mq,
		Tracer: tracer,
		// TracerClose: cleanup,
		Monitoring: monitoring,
		DB:         db,
	}, nil
}

func NewLoggerService(log logger.Logger, mq mq.MQ) (*Service, error) {
	return &Service{
		Log: log,
		MQ:  mq,
	}, nil
}

var FullSet = wire.NewSet(NewFullService, InitLogger, InitStore, InitTracer, InitMonitoring, InitMQ)
var LoggerSet = wire.NewSet(NewLoggerService, InitLogger, InitMQ)

func InitializeFullService(ctx context.Context) (*Service, error) {
	panic(wire.Build(FullSet))
}

func InitializeLoggerService(ctx context.Context) (*Service, error) {
	panic(wire.Build(LoggerSet))
}
