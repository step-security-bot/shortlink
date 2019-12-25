package api

import (
	"context"

	"github.com/batazor/shortlink/internal/logger"
	api_type "github.com/batazor/shortlink/pkg/api/type"
)

// API - general describe of API
type API interface { // nolint unused
	Run(ctx context.Context, config api_type.Config, log logger.Logger) error
}

type Server struct{} // nolint unused
