// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package notify_di

import (
	"context"
	"github.com/google/wire"
	"github.com/shortlink-org/shortlink/internal/di"
	"github.com/shortlink-org/shortlink/internal/di/pkg/autoMaxPro"
	"github.com/shortlink-org/shortlink/internal/di/pkg/config"
	"github.com/shortlink-org/shortlink/internal/di/pkg/context"
	"github.com/shortlink-org/shortlink/internal/di/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/di/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/di/pkg/profiling"
	"github.com/shortlink-org/shortlink/internal/di/pkg/traicing"
	"github.com/shortlink-org/shortlink/internal/pkg/logger"
	"github.com/shortlink-org/shortlink/internal/pkg/monitoring"
	"github.com/shortlink-org/shortlink/internal/pkg/mq"
	"github.com/shortlink-org/shortlink/internal/services/notify/application"
	"github.com/shortlink-org/shortlink/internal/services/notify/infrastructure/slack"
	"github.com/shortlink-org/shortlink/internal/services/notify/infrastructure/smtp"
	"github.com/shortlink-org/shortlink/internal/services/notify/infrastructure/telegram"
	"go.opentelemetry.io/otel/trace"
)

// Injectors from wire.go:

func InitializeFullBotService() (*Service, func(), error) {
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
	monitoringMonitoring, cleanup3, err := monitoring.New(context, logger)
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
	pprofEndpoint, err := profiling.New(logger)
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
	bot := InitSlack(context)
	telegramBot := InitTelegram(context)
	smtpBot := InitSMTP(context)
	dataBus, cleanup6, err := mq_di.New(context, logger)
	if err != nil {
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	applicationBot, err := NewBotApplication(context, logger, dataBus)
	if err != nil {
		cleanup6()
		cleanup5()
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	service, err := NewBotService(logger, configConfig, monitoringMonitoring, tracerProvider, pprofEndpoint, autoMaxProAutoMaxPro, bot, telegramBot, smtpBot, applicationBot)
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

// wire.go:

// Service - heplers
type Service struct {
	// Common
	Log    logger.Logger
	Config *config.Config

	// Observability
	Tracer        trace.TracerProvider
	Monitoring    *monitoring.Monitoring
	PprofEndpoint profiling.PprofEndpoint
	AutoMaxPro    autoMaxPro.AutoMaxPro

	// Bot
	slack    *slack.Bot
	telegram *telegram.Bot
	smtp     *smtp.Bot

	// Application
	botService *application.Bot
}

// InitSlack - Init slack bot
func InitSlack(ctx2 context.Context) *slack.Bot {
	slackBot := &slack.Bot{}
	if err := slackBot.Init(); err != nil {
		return nil
	}

	return slackBot
}

// InitTelegram - Init telegram bot
func InitTelegram(ctx2 context.Context) *telegram.Bot {
	telegramBot := &telegram.Bot{}
	if err := telegramBot.Init(); err != nil {
		return nil
	}

	return telegramBot
}

// InitSMTP - Init SMTP bot
func InitSMTP(ctx2 context.Context) *smtp.Bot {
	smtpBot := &smtp.Bot{}
	if err := smtpBot.Init(); err != nil {
		return nil
	}

	return smtpBot
}

// FullBotService ======================================================================================================
var NotifySet = wire.NewSet(di.DefaultSet, mq_di.New, InitSlack,
	InitTelegram,
	InitSMTP,

	NewBotApplication,

	NewBotService,
)

func NewBotApplication(ctx2 context.Context, logger2 logger.Logger, mq2 *mq.DataBus) (*application.Bot, error) {
	bot, err := application.New(mq2, logger2)
	if err != nil {
		return nil, err
	}
	bot.Use(ctx2)

	return bot, nil
}

func NewBotService(

	log logger.Logger, config2 *config.Config, monitoring2 *monitoring.Monitoring,
	tracer trace.TracerProvider,
	pprofHTTP profiling.PprofEndpoint,
	autoMaxProcsOption autoMaxPro.AutoMaxPro, slack2 *slack.Bot, telegram2 *telegram.Bot, smtp2 *smtp.Bot,

	bot *application.Bot,
) (*Service, error) {
	return &Service{

		Log:    log,
		Config: config2,

		Tracer:        tracer,
		Monitoring:    monitoring2,
		PprofEndpoint: pprofHTTP,
		AutoMaxPro:    autoMaxProcsOption,

		slack:    slack2,
		telegram: telegram2,
		smtp:     smtp2,

		botService: bot,
	}, nil
}
