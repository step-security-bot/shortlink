// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bff_web_di

import (
	"context"
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/di/pkg/monitoring"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	http2 "github.com/shortlink-org/shortlink/internal/services/bff-web/infrastructure/http"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

// Injectors from wire.go:

func InitializeBFFWebService() (*BFFWebService, func(), error) {
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
	tracerProvider, cleanup3, err := traicing_di.New(context, logger)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	serveMux, err := monitoring.New(logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint, err := profiling.New(logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	autoMaxProAutoMaxPro, cleanup4, err := autoMaxPro.New(logger)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server, err := BFFWebAPIService(context, logger, tracerProvider)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	bffWebService := NewBFFWebService(context, logger, configConfig, tracerProvider, serveMux, pprofEndpoint, autoMaxProAutoMaxPro, server)
	return bffWebService, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type BFFWebService struct {
	// Common
	Logger logger.Logger
	Config *config.Config

	// Observability
	Tracer        *trace.TracerProvider
	Monitoring    *http.ServeMux
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Delivery
	httpAPIServer *http2.Server
}

// BFFWebService =======================================================================================================
var BFFWebServiceSet = wire.NewSet(di.DefaultSet, BFFWebAPIService,

	NewBFFWebService,
)

func BFFWebAPIService(ctx2 context.Context, logger2 logger.Logger,

	tracer *trace.TracerProvider,
) (*http2.Server, error) {

	API := http2.Server{}
	apiService, err := API.Run(ctx2, logger2, tracer)
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

func NewBFFWebService(ctx2 context.Context, logger2 logger.Logger, config2 *config.Config,

	tracer *trace.TracerProvider, monitoring2 *http.ServeMux,
	pprofEndpoint profiling.PprofEndpoint, autoMaxPro2 autoMaxPro.AutoMaxPro,

	httpAPIServer *http2.Server,
) *BFFWebService {
	return &BFFWebService{

		Logger: logger2,
		Config: config2,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofEndpoint,
		AutoMaxPro:    autoMaxPro2,

		httpAPIServer: httpAPIServer,
	}
}
