//+build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"context"
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"github.com/batazor/shortlink/internal/di/internal/autoMaxPro"
	"github.com/batazor/shortlink/internal/di/internal/config"
	"github.com/batazor/shortlink/internal/di/internal/monitoring"
	mq_di "github.com/batazor/shortlink/internal/di/internal/mq"
	"github.com/batazor/shortlink/internal/di/internal/profiling"
	"github.com/batazor/shortlink/internal/di/internal/sentry"
	"github.com/batazor/shortlink/internal/di/internal/store"
	"github.com/batazor/shortlink/internal/pkg/db"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/batazor/shortlink/internal/pkg/mq"
	"github.com/batazor/shortlink/pkg/rpc"
)

type ServiceAPI struct {
	Service
}

// APIService ==========================================================================================================
var APISet = wire.NewSet(
	DefaultSet,
	store.New,
	sentry.New,
	monitoring.New,
	profiling.New,
	mq_di.New,
	rpc.InitServer,
	rpc.InitClient,
	NewAPIService,
)

func NewAPIService(
	ctx context.Context,
	cfg *config.Config,
	log logger.Logger,
	mq mq.MQ,
	sentryHandler *sentryhttp.Handler,
	monitoring *http.ServeMux,
	tracer *opentracing.Tracer,
	db *db.Store,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro,
	clientRPC *grpc.ClientConn,
) (*ServiceAPI, error) {
	return &ServiceAPI{
		Service: Service{
			Ctx:           ctx,
			Log:           log,
			MQ:            mq,
			Tracer:        tracer,
			Monitoring:    monitoring,
			Sentry:        sentryHandler,
			DB:            db,
			PprofEndpoint: pprofHTTP,
			ClientRPC:     clientRPC,
		},
	}, nil
}

func InitializeAPIService() (*ServiceAPI, func(), error) {
	panic(wire.Build(APISet))
}
