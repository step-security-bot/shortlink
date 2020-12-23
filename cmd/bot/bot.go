/*
Bot application
*/
package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"google.golang.org/protobuf/proto"

	"github.com/batazor/shortlink/internal/api/domain/link"
	"github.com/batazor/shortlink/internal/bot"
	bot_type "github.com/batazor/shortlink/internal/bot/type"
	"github.com/batazor/shortlink/internal/di"
	"github.com/batazor/shortlink/internal/logger/field"
	"github.com/batazor/shortlink/internal/mq/query"
	"github.com/batazor/shortlink/internal/notify"
)

func main() {
	viper.SetDefault("SERVICE_NAME", "bot")

	// Init a new service
	s, cleanup, err := di.InitializeBotService()
	if err != nil { // TODO: use as helpers
		var typeErr *net.OpError
		if errors.As(err, &typeErr) {
			panic(fmt.Errorf("address %s already in use. Set GRPC_SERVER_PORT environment", typeErr.Addr.String()))
		}

		panic(err)
	}

	getEventNewLink := query.Response{
		Chan: make(chan []byte),
	}

	// Run bot
	b := bot.Bot{}
	b.Use(s.Ctx)

	go func() {
		if s.MQ != nil {
			if err := s.MQ.Subscribe(getEventNewLink); err != nil {
				s.Log.Error(err.Error())
			}
		}
	}()

	go func() {
		for {
			msg := <-getEventNewLink.Chan

			// []byte to link.Link
			myLink := &link.Link{}
			if err := proto.Unmarshal(msg, myLink); err != nil {
				s.Log.Error(fmt.Sprintf("Error unmarsharing event new link: %s", err.Error()))
				continue
			}

			s.Log.Info("Get new LINK", field.Fields{"url": myLink.Url})
			notify.Publish(s.Ctx, bot_type.METHOD_NEW_LINK, myLink, nil)
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			s.Log.Error(r.(string))
		}
	}()

	// Handle SIGINT and SIGTERM.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// Stop the service gracefully.
	cleanup()
}
