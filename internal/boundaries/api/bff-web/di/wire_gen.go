// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bff_web_di

import (
	"context"
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/internal/boundaries/api/bff-web/infrastructure/http"
	v1_2 "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/cqrs/link/v1"
	"github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/link/v1"
	v1_3 "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/sitemap/v1"
	v1_4 "github.com/shortlink-org/shortlink/internal/boundaries/link/metadata/infrastructure/rpc/metadata/v1"
	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/internal/pkg/i18n"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/internal/pkg/rpc"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/text/message"
	"google.golang.org/grpc"
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
	monitoringMonitoring, cleanup4, err := monitoring.New(context, logger)
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
	printer := i18n.New(context)
	server, err := rpc.InitServer(context, logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	clientConn, cleanup6, err := rpc.InitClient(context, logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	linkServiceClient, err := NewLinkRPCClient(clientConn)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	linkCommandServiceClient, err := NewLinkCommandRPCClient(clientConn)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	linkQueryServiceClient, err := NewLinkQueryRPCClient(clientConn)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	sitemapServiceClient, err := NewSitemapServiceClient(clientConn)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	httpServer, err := NewAPIApplication(context, printer, logger, configConfig, tracerProvider, monitoringMonitoring, pprofEndpoint, autoMaxProAutoMaxPro, server, linkServiceClient, linkCommandServiceClient, linkQueryServiceClient, sitemapServiceClient)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	bffWebService := NewBFFWebService(context, logger, configConfig, tracerProvider, monitoringMonitoring, pprofEndpoint, autoMaxProAutoMaxPro, httpServer)
	return bffWebService, func() {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type BFFWebService struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Delivery
	httpAPIServer *http.Server

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro
}

// BFFWebService =======================================================================================================
var BFFWebServiceSet = wire.NewSet(di.DefaultSet, rpc.InitServer, rpc.InitClient, NewLinkRPCClient,
	NewLinkCommandRPCClient,
	NewLinkQueryRPCClient,
	NewSitemapServiceClient,
	NewMetadataRPCClient,

	NewAPIApplication,
	NewBFFWebService,
)

func NewLinkRPCClient(runRPCClient *grpc.ClientConn) (v1.LinkServiceClient, error) {
	LinkServiceClient := v1.NewLinkServiceClient(runRPCClient)
	return LinkServiceClient, nil
}

func NewLinkCommandRPCClient(runRPCClient *grpc.ClientConn) (v1_2.LinkCommandServiceClient, error) {
	LinkCommandRPCClient := v1_2.NewLinkCommandServiceClient(runRPCClient)
	return LinkCommandRPCClient, nil
}

func NewLinkQueryRPCClient(runRPCClient *grpc.ClientConn) (v1_2.LinkQueryServiceClient, error) {
	LinkQueryRPCClient := v1_2.NewLinkQueryServiceClient(runRPCClient)
	return LinkQueryRPCClient, nil
}

func NewSitemapServiceClient(runRPCClient *grpc.ClientConn) (v1_3.SitemapServiceClient, error) {
	sitemapRPCClient := v1_3.NewSitemapServiceClient(runRPCClient)
	return sitemapRPCClient, nil
}

func NewMetadataRPCClient(runRPCClient *grpc.ClientConn) (v1_4.MetadataServiceClient, error) {
	metadataRPCClient := v1_4.NewMetadataServiceClient(runRPCClient)
	return metadataRPCClient, nil
}

func NewAPIApplication(ctx2 context.Context, i18n2 *message.Printer,
	log logger.Logger, config2 *config.Config,

	tracer trace.TracerProvider, monitoring2 *monitoring.Monitoring,
	pprofEndpoint profiling.PprofEndpoint, autoMaxPro2 autoMaxPro.AutoMaxPro,

	rpcServer *rpc.Server,
	link_rpc v1.LinkServiceClient,
	link_command v1_2.LinkCommandServiceClient,
	link_query v1_2.LinkQueryServiceClient,
	sitemap_rpc v1_3.SitemapServiceClient,
) (*http.Server, error) {
	apiService, err := http.New(http.Config{

		Ctx:    ctx2,
		I18n:   i18n2,
		Log:    log,
		Config: config2,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofEndpoint,
		AutoMaxPro:    autoMaxPro2,

		RpcServer: rpcServer,

		Link_rpc:     link_rpc,
		Link_command: link_command,
		Link_query:   link_query,
		Sitemap_rpc:  sitemap_rpc,
	})
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

func NewBFFWebService(ctx2 context.Context,

	log logger.Logger, config2 *config.Config,

	tracer trace.TracerProvider, monitoring2 *monitoring.Monitoring,
	pprofEndpoint profiling.PprofEndpoint, autoMaxPro2 autoMaxPro.AutoMaxPro,

	httpAPIServer *http.Server,
) *BFFWebService {
	return &BFFWebService{

		Log:    log,
		Config: config2,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofEndpoint,
		AutoMaxPro:    autoMaxPro2,

		httpAPIServer: httpAPIServer,
	}
}
