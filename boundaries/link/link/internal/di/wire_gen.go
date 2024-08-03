// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package link_di

import (
	"github.com/authzed/authzed-go/v1"
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/mq"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/cqrs/cqs"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/cqrs/query"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/repository/crud"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/rpc/cqrs/link/v1"
	v1_2 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/rpc/link/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/rpc/run"
	v1_3 "github.com/shortlink-org/shortlink/boundaries/link/link/internal/infrastructure/rpc/sitemap/v1"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/usecases/link"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/usecases/link_cqrs"
	"github.com/shortlink-org/shortlink/boundaries/link/link/internal/usecases/sitemap"
	"github.com/shortlink-org/shortlink/pkg/cache"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/context"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/permission"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/store"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/mq"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"
	"go.opentelemetry.io/otel/trace"
)

// Injectors from wire.go:

func InitializeLinkService() (*LinkService, func(), error) {
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
	autoMaxProAutoMaxPro, cleanup3, err := autoMaxPro.New(logger)
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
	monitoringMonitoring, cleanup5, err := monitoring.New(context, logger, tracerProvider)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	pprofEndpoint, err := profiling.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	client, err := permission.New(context, logger, tracerProvider, monitoringMonitoring)
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
	db, err := store.New(context, logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	cacheCache, err := cache.New(context, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	crudStore, err := crud.New(context, logger, db, cacheCache)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	uc, err := NewLinkApplication(logger, mq, crudStore, client)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	cqsStore, err := cqs.New(context, logger, db, cacheCache)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	queryStore, err := query.New(context, logger, db, cacheCache)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := link_cqrs.New(logger, cqsStore, queryStore)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	sitemapService, err := sitemap.New(logger, mq)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	event, err := api_mq.New(mq, logger, uc)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	server, err := rpc.InitServer(context, logger, tracerProvider, monitoringMonitoring)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	linkRPC, err := v1.New(server, service, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	v1LinkRPC, err := v1_2.New(server, uc, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	response, err := NewRunRPCServer(server, linkRPC, v1LinkRPC)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	v1Sitemap, err := v1_3.New(server, sitemapService, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	linkService, err := NewLinkService(logger, configConfig, autoMaxProAutoMaxPro, monitoringMonitoring, tracerProvider, pprofEndpoint, client, uc, service, sitemapService, event, response, v1LinkRPC, linkRPC, v1Sitemap, crudStore, cqsStore, queryStore)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return linkService, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type LinkService struct {
	// Common
	Log        logger.Logger
	Config     *config.Config
	AutoMaxPro autoMaxPro.AutoMaxPro

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint

	// Security
	authPermission *authzed.Client

	// Delivery
	linkMQ            *api_mq.Event
	run               *run.Response
	linkRPCServer     *v1_2.LinkRPC
	linkCQRSRPCServer *v1.LinkRPC
	sitemapRPCServer  *v1_3.Sitemap

	// Application
	linkService     *link.UC
	linkCQRSService *link_cqrs.Service
	sitemapService  *sitemap.Service

	// Repository
	linkStore *crud.Store

	// CQRS
	cqsStore   *cqs.Store
	queryStore *query.Store
}

// LinkService =========================================================================================================
var LinkSet = wire.NewSet(di.DefaultSet, mq_di.New, rpc.InitServer, rpc.InitClient, store.New, api_mq.New, v1_2.New, v1.New, v1_3.New, NewRunRPCServer, v1_2.NewLinkServiceClient, NewLinkApplication, link_cqrs.New, sitemap.New, crud.New, cqs.New, query.New, NewLinkService)

func NewLinkApplication(log logger.Logger, mq2 mq.MQ, store2 *crud.Store, authPermission *authzed.Client) (*link.UC, error) {
	linkService, err := link.New(log, mq2, nil, store2, authPermission)
	if err != nil {
		return nil, err
	}

	return linkService, nil
}

// TODO: refactoring. maybe drop this function
func NewRunRPCServer(runRPCServer *rpc.Server, _ *v1.LinkRPC, _ *v1_2.LinkRPC) (*run.Response, error) {
	return run.Run(runRPCServer)
}

func NewLinkService(

	log logger.Logger, config2 *config.Config,
	autoMaxProcsOption autoMaxPro.AutoMaxPro, monitoring2 *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,

	authPermission *authzed.Client,

	linkService *link.UC,
	linkCQRSService *link_cqrs.Service,
	sitemapService *sitemap.Service,

	linkMQ *api_mq.Event, run2 *run.Response,
	linkRPCServer *v1_2.LinkRPC,
	linkCQRSRPCServer *v1.LinkRPC,
	sitemapRPCServer *v1_3.Sitemap,

	linkStore *crud.Store,

	cqsStore *cqs.Store,
	queryStore *query.Store,
) (*LinkService, error) {
	return &LinkService{

		Log:        log,
		Config:     config2,
		AutoMaxPro: autoMaxProcsOption,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofHTTP,

		authPermission: authPermission,

		linkService:     linkService,
		linkCQRSService: linkCQRSService,
		sitemapService:  sitemapService,

		run:               run2,
		linkRPCServer:     linkRPCServer,
		linkCQRSRPCServer: linkCQRSRPCServer,
		sitemapRPCServer:  sitemapRPCServer,
		linkMQ:            linkMQ,

		linkStore: linkStore,

		cqsStore:   cqsStore,
		queryStore: queryStore,
	}, nil
}
