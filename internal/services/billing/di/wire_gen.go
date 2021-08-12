// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package billing_di

import (
	"context"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/eventsourcing/store"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq/v1"
	"github.com/batazor/shortlink/internal/services/billing/application/account"
	"github.com/batazor/shortlink/internal/services/billing/application/order"
	"github.com/batazor/shortlink/internal/services/billing/application/payment"
	"github.com/batazor/shortlink/internal/services/billing/application/tariff"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/api/http"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/api/rpc/order/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/api/rpc/payment/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/api/rpc/tariff/v1"
	"github.com/batazor/shortlink/internal/services/billing/infrastructure/store"
	"github.com/batazor/shortlink/pkg/rpc"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
)

// Injectors from di.go:

func InitializeBillingService(ctx context.Context, runRPCClient *grpc.ClientConn, runRPCServer *rpc.RPCServer, log logger.Logger, db2 *db.Store, mq v1.MQ, tracer *opentracing.Tracer) (*BillingService, func(), error) {
	billingStore, err := NewBillingStore(ctx, log, db2)
	if err != nil {
		return nil, nil, err
	}
	accountService, err := NewAccountApplication(log, billingStore)
	if err != nil {
		return nil, nil, err
	}
	orderService, err := NewOrderApplication(log, billingStore)
	if err != nil {
		return nil, nil, err
	}
	paymentService, err := NewPaymentApplication(log, billingStore)
	if err != nil {
		return nil, nil, err
	}
	tariffService, err := NewTariffApplication(log, billingStore)
	if err != nil {
		return nil, nil, err
	}
	server, err := NewBillingAPIServer(ctx, log, tracer, runRPCServer, db2, accountService, orderService, paymentService, tariffService)
	if err != nil {
		return nil, nil, err
	}
	billingService, err := NewBillingService(log, server)
	if err != nil {
		return nil, nil, err
	}
	return billingService, func() {
	}, nil
}

// di.go:

type BillingService struct {
	Logger logger.Logger

	// Delivery
	httpAPIServer    *api.Server
	orderRPCServer   *order_rpc.Order
	paymentRPCServer *payment_rpc.Payment
	tariffRPCServer  *tariff_rpc.Tariff

	// Repository
	accountRepository    *billing_store.AccountRepository
	tariffRepository     *billing_store.TariffRepository
	eventStoreRepository *event_store.EventStore
}

// BillingService ======================================================================================================
var BillingSet = wire.NewSet(

	NewBillingAPIServer,

	NewBillingStore,

	NewTariffApplication,
	NewAccountApplication,
	NewOrderApplication,
	NewPaymentApplication,

	NewBillingService,
)

func NewBillingStore(ctx context.Context, logger2 logger.Logger, db2 *db.Store) (*billing_store.BillingStore, error) {
	store := &billing_store.BillingStore{}
	billingStore, err := store.Use(ctx, logger2, db2)
	if err != nil {
		return nil, err
	}

	return billingStore, nil
}

func NewAccountApplication(logger2 logger.Logger, store *billing_store.BillingStore) (*account_application.AccountService, error) {
	accountService, err := account_application.New(logger2, store.Account)
	if err != nil {
		return nil, err
	}

	return accountService, nil
}

func NewOrderApplication(logger2 logger.Logger, store *billing_store.BillingStore) (*order_application.OrderService, error) {
	orderService, err := order_application.New(logger2, store.EventStore)
	if err != nil {
		return nil, err
	}

	return orderService, nil
}

func NewPaymentApplication(logger2 logger.Logger, store *billing_store.BillingStore) (*payment_application.PaymentService, error) {
	paymentService, err := payment_application.New(logger2, store.EventStore)
	if err != nil {
		return nil, err
	}

	return paymentService, nil
}

func NewTariffApplication(logger2 logger.Logger, store *billing_store.BillingStore) (*tariff_application.TariffService, error) {
	tariffService, err := tariff_application.New(logger2, store.Tariff)
	if err != nil {
		return nil, err
	}

	return tariffService, nil
}

func NewBillingAPIServer(
	ctx context.Context, logger2 logger.Logger,

	tracer *opentracing.Tracer,
	rpcServer *rpc.RPCServer, db2 *db.Store,

	accountService *account_application.AccountService,
	orderService *order_application.OrderService,
	paymentService *payment_application.PaymentService,
	tariffService *tariff_application.TariffService,
) (*api.Server, error) {

	API := api.Server{}

	apiService, err := API.Use(
		ctx, db2, logger2, tracer,

		accountService,
		orderService,
		paymentService,
		tariffService,
	)
	if err != nil {
		return nil, err
	}

	return apiService, nil
}

func NewBillingService(
	log logger.Logger,

	httpAPIServer *api.Server,
) (*BillingService, error) {
	return &BillingService{
		Logger: log,

		httpAPIServer: httpAPIServer,
	}, nil
}
