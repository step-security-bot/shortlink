// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package di

import (
	"context"
	"fmt"
	"github.com/batazor/shortlink/internal/logger"
	"github.com/batazor/shortlink/internal/mq"
	"github.com/batazor/shortlink/internal/store"
	"github.com/batazor/shortlink/internal/traicing"
	"github.com/getsentry/sentry-go"
	"github.com/getsentry/sentry-go/http"
	"github.com/google/wire"
	"github.com/heptiolabs/healthcheck"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"
	"go.uber.org/automaxprocs/maxprocs"
	"net/http"
	"net/http/pprof"
	"time"
)

// Injectors from wire.go:

func InitializeFullService(ctx context.Context) (*Service, func(), error) {
	logger, cleanup, err := InitLogger(ctx)
	if err != nil {
		return nil, nil, err
	}
	mq, cleanup2, err := InitMQ(ctx, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	handler, cleanup3, err := InitSentry()
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux := InitMonitoring(handler)
	tracer, cleanup4, err := InitTracer(ctx, logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	db, cleanup5, err := InitStore(ctx, logger)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint := InitProfiling()
	diDiAutoMaxPro, cleanup6, err := InitAutoMaxProcs(logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewFullService(logger, mq, serveMux, tracer, db, pprofEndpoint, handler, diDiAutoMaxPro)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

func InitializeLoggerService(ctx context.Context) (*Service, func(), error) {
	logger, cleanup, err := InitLogger(ctx)
	if err != nil {
		return nil, nil, err
	}
	mq, cleanup2, err := InitMQ(ctx, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	diDiAutoMaxPro, cleanup3, err := InitAutoMaxProcs(logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewLoggerService(logger, mq, diDiAutoMaxPro)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

// Service - heplers
type Service struct {
	Log    logger.Logger
	Tracer opentracing.Tracer
	// TracerClose func()
	Sentry        *sentryhttp.Handler
	DB            store.DB
	MQ            mq.MQ
	Monitoring    *http.ServeMux
	PprofEndpoint PprofEndpoint
}

type PprofEndpoint *http.ServeMux

type diAutoMaxPro *string

// InitAutoMaxProcs - Automatically set GOMAXPROCS to match Linux container CPU quota
func InitAutoMaxProcs(log logger.Logger) (diAutoMaxPro, func(), error) {
	undo, err := maxprocs.Set(maxprocs.Logger(func(s string, args ...interface{}) {
		log.Info(fmt.Sprintf(s, args))
	}))
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		undo()
	}

	return nil, cleanup, nil
}

// InitStore return store
func InitStore(ctx context.Context, log logger.Logger) (store.DB, func(), error) {
	var st store.Store
	db := st.Use(ctx, log)

	cleanup := func() {
		if err := db.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return db, cleanup, nil
}

func InitLogger(ctx context.Context) (logger.Logger, func(), error) {
	viper.SetDefault("LOG_LEVEL", logger.INFO_LEVEL)
	viper.SetDefault("LOG_TIME_FORMAT", time.RFC3339Nano)

	conf := logger.Configuration{
		Level:      viper.GetInt("LOG_LEVEL"),
		TimeFormat: viper.GetString("LOG_TIME_FORMAT"),
	}

	log, err := logger.NewLogger(logger.Zap, conf)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {

		_ = log.Close()
	}

	return log, cleanup, nil
}

func InitTracer(ctx context.Context, log logger.Logger) (opentracing.Tracer, func(), error) {
	viper.SetDefault("TRACER_SERVICE_NAME", "ShortLink")
	viper.SetDefault("TRACER_URI", "localhost:6831")

	config := traicing.Config{
		ServiceName: viper.GetString("TRACER_SERVICE_NAME"),
		URI:         viper.GetString("TRACER_URI"),
	}

	tracer, tracerClose, err := traicing.Init(config)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		if err := tracerClose.Close(); err != nil {
			log.Error(err.Error())
		}
	}

	return tracer, cleanup, nil
}

func InitMQ(ctx context.Context, log logger.Logger) (mq.MQ, func(), error) {
	viper.SetDefault("MQ_ENABLED", "false")

	if viper.GetBool("MQ_ENABLED") {
		var service mq.DataBus
		dataBus := service.Use(ctx, log)

		cleanup := func() {
			if err := dataBus.Close(); err != nil {
				log.Error(err.Error())
			}
		}

		return dataBus, cleanup, nil
	}

	return nil, nil, nil
}

func InitMonitoring(sentryHandler *sentryhttp.Handler) *http.ServeMux {

	registry := prometheus.NewRegistry()

	health := healthcheck.NewMetricsHandler(registry, "common")

	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	commonMux := http.NewServeMux()

	commonMux.Handle("/metrics", sentryHandler.Handle(promhttp.HandlerFor(registry, promhttp.HandlerOpts{})))

	commonMux.HandleFunc("/live", sentryHandler.HandleFunc(health.LiveEndpoint))

	commonMux.HandleFunc("/ready", sentryHandler.HandleFunc(health.ReadyEndpoint))

	return commonMux
}

func InitProfiling() PprofEndpoint {

	pprofMux := http.NewServeMux()

	pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
	pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	return pprofMux
}

func InitSentry() (*sentryhttp.Handler, func(), error) {
	DSN := viper.GetString("SENTRY_DSN")

	if DSN != "" {
		return nil, nil, nil
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("DSN"),
	})
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		sentry.Flush(time.Second * 5)
		sentry.Recover()
	}

	sentryHandler := sentryhttp.New(sentryhttp.Options{})

	return sentryHandler, cleanup, nil
}

// Default =============================================================================================================
var DefaultSet = wire.NewSet(InitAutoMaxProcs, InitLogger, InitTracer)

// FullService =========================================================================================================
var FullSet = wire.NewSet(DefaultSet, NewFullService, InitStore, InitMonitoring, InitProfiling, InitMQ, InitSentry)

func NewFullService(log logger.Logger, mq2 mq.MQ, monitoring *http.ServeMux, tracer opentracing.Tracer, db store.DB, pprofHTTP PprofEndpoint, sentryHandler *sentryhttp.Handler, autoMaxProcsOption diAutoMaxPro) (*Service, error) {
	return &Service{
		Log:    log,
		MQ:     mq2,
		Tracer: tracer,

		Monitoring:    monitoring,
		DB:            db,
		PprofEndpoint: pprofHTTP,
		Sentry:        sentryHandler,
	}, nil
}

// LoggerService =======================================================================================================
var LoggerSet = wire.NewSet(DefaultSet, NewLoggerService, InitMQ)

func NewLoggerService(log logger.Logger, mq2 mq.MQ, autoMaxProcsOption diAutoMaxPro) (*Service, error) {
	return &Service{
		Log: log,
		MQ:  mq2,
	}, nil
}
