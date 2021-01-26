// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package di

import (
	"context"
	"github.com/batazor/shortlink/internal/services/bot/infrastructure/slack"
	"github.com/batazor/shortlink/internal/services/bot/infrastructure/smtp"
	"github.com/batazor/shortlink/internal/services/bot/infrastructure/telegram"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeFullBotService() (*Service, func(), error) {
	context, cleanup, err := NewContext()
	if err != nil {
		return nil, nil, err
	}
	bot := InitSlack(context)
	telegramBot := InitTelegram(context)
	smtpBot := InitSMTP(context)
	service, err := NewBotService(bot, telegramBot, smtpBot)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return service, func() {
		cleanup()
	}, nil
}

// wire.go:

// Service - heplers
type Service struct {
	slack    *slack.Bot
	telegram *telegram.Bot
	smtp     *smtp.Bot
}

// Context =============================================================================================================
func NewContext() (context.Context, func(), error) {
	ctx := context.Background()

	cb := func() {
		ctx.Done()
	}

	return ctx, cb, nil
}

// InitSlack - Init slack bot
func InitSlack(ctx context.Context) *slack.Bot {
	slackBot := &slack.Bot{}
	if err := slackBot.Init(); err != nil {
		return nil
	}

	return slackBot
}

// InitTelegram - Init telegram bot
func InitTelegram(ctx context.Context) *telegram.Bot {
	telegramBot := &telegram.Bot{}
	if err := telegramBot.Init(); err != nil {
		return nil
	}

	return telegramBot
}

// InitSMTP - Init SMTP bot
func InitSMTP(ctx context.Context) *smtp.Bot {
	smtpBot := &smtp.Bot{}
	if err := smtpBot.Init(); err != nil {
		return nil
	}

	return smtpBot
}

// FullBotService ======================================================================================================
var FullBotSet = wire.NewSet(NewContext, InitSlack, InitTelegram, InitSMTP, NewBotService)

func NewBotService(slack2 *slack.Bot, telegram2 *telegram.Bot, smtp2 *smtp.Bot) (*Service, error) {
	return &Service{slack2, telegram2, smtp2}, nil
}
