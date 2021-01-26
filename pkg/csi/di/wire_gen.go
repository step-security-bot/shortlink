// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"context"
	"github.com/batazor/shortlink/internal/pkg/logger"
	"github.com/google/wire"
	"github.com/spf13/viper"
	"time"
)

// Injectors from wire.go:

func InitializeSCIDriver() (*Service, func(), error) {
	context, cleanup, err := NewContext()
	if err != nil {
		return nil, nil, err
	}
	logger, cleanup2, err := InitLogger(context)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	service, err := NewSCIDriver(logger, context)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

// Service - heplers
type Service struct {
	Ctx context.Context
	Log logger.Logger
}

// Context =============================================================================================================
func NewContext() (context.Context, func(), error) {
	ctx := context.Background()

	cb := func() {
		ctx.Done()
	}

	return ctx, cb, nil
}

// Logger ==============================================================================================================
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

// CSI =================================================================================================================
var FullBotSet = wire.NewSet(NewContext, InitLogger, NewSCIDriver)

func NewSCIDriver(log logger.Logger, ctx context.Context) (*Service, error) {
	return &Service{
		Ctx: ctx,
		Log: log,
	}, nil
}
