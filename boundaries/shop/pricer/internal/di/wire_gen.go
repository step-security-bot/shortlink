// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/application"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/infrastructure"
	"github.com/shortlink-org/shortlink/boundaries/shop/pricer/internal/interfaces/cli"
	"github.com/shortlink-org/shortlink/pkg/di"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/config"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/context"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/pkg/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"
	"github.com/shortlink-org/shortlink/pkg/rpc"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel/trace"
	"log"
)

// Injectors from wire.go:

func InitializePricerService() (*PricerService, func(), error) {
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
	discountPolicy := newDiscountPolicy()
	taxPolicy := newTaxPolicy()
	v, err := newPolicyNames()
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	cartService := application.NewCartService(discountPolicy, taxPolicy, v)
	pricerService, err := NewPricerService(logger, configConfig, autoMaxProAutoMaxPro, monitoringMonitoring, tracerProvider, pprofEndpoint, cartService)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return pricerService, func() {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

type PricerService struct {
	// Common
	Log        logger.Logger
	Config     *config.Config
	AutoMaxPro autoMaxPro.AutoMaxPro

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint

	// Application
	CartService *application.CartService

	// CLI
	CLIHandler *cli.CLIHandler
}

// PricerService =======================================================================================================
var PricerSet = wire.NewSet(di.DefaultSet, rpc.InitServer, newDiscountPolicy,
	newTaxPolicy,
	newPolicyNames, application.NewCartService, NewPricerService,
)

// newDiscountPolicy creates a new DiscountPolicy
func newDiscountPolicy() application.DiscountPolicy {
	discountPolicyPath := viper.GetString("policies.discounts")
	discountQuery := viper.GetString("queries.discounts")

	discountEvaluator, err := infrastructure.NewOPAEvaluator(discountPolicyPath, discountQuery)
	if err != nil {
		log.Fatalf("Failed to initialize Discount Policy Evaluator: %v", err)
	}

	return discountEvaluator
}

// newTaxPolicy creates a new TaxPolicy
func newTaxPolicy() application.TaxPolicy {
	taxPolicyPath := viper.GetString("policies.taxes")
	taxQuery := viper.GetString("queries.taxes")

	taxEvaluator, err := infrastructure.NewOPAEvaluator(taxPolicyPath, taxQuery)
	if err != nil {
		log.Fatalf("Failed to initialize Tax Policy Evaluator: %v", err)
	}

	return taxEvaluator
}

// newPolicyNames retrieves policy names
func newPolicyNames() ([]string, error) {
	discountPolicyPath := viper.GetString("policies.discounts")
	taxPolicyPath := viper.GetString("policies.taxes")

	return infrastructure.GetPolicyNames(discountPolicyPath, taxPolicyPath)
}

func NewPricerService(log2 logger.Logger, config2 *config.Config,
	autoMaxProcsOption autoMaxPro.AutoMaxPro, monitoring2 *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,

	cartService *application.CartService,
) (*PricerService, error) {
	return &PricerService{

		Log:        log2,
		Config:     config2,
		AutoMaxPro: autoMaxProcsOption,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofHTTP,

		CartService: cartService,
	}, nil
}
