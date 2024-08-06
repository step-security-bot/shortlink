package temporal

import (
	"github.com/uber-go/tally/v4/prometheus"
	"go.temporal.io/sdk/client"

	"github.com/shortlink-org/shortlink/pkg/logger"
	"github.com/shortlink-org/shortlink/pkg/observability/monitoring"

	sdktally "go.temporal.io/sdk/contrib/tally"
)

func New(log logger.Logger, monitoring *monitoring.Monitoring) (client.Client, error) {
	metrics, err := newPrometheusScope(prometheus.Configuration{
		ListenAddress: "0.0.0.0:9090",
		TimerType:     "histogram",
	}, monitoring, log)
	if err != nil {
		return nil, err
	}

	// Create the client object just once per process
	c, err := client.Dial(client.Options{
		MetricsHandler: sdktally.NewMetricsHandler(metrics),
	})
	if err != nil {
		return nil, err
	}

	return c, nil
}
