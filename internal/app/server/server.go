package server

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/igor-benko/pow-tcp-server/internal/config"
	"github.com/igor-benko/pow-tcp-server/internal/delivery/tcp"
	"github.com/igor-benko/pow-tcp-server/internal/domain"
	"github.com/igor-benko/pow-tcp-server/internal/repository/storage/memory"
	"github.com/igor-benko/pow-tcp-server/internal/service/quote"
	"github.com/igor-benko/pow-tcp-server/pkg/pow"
	"github.com/rs/zerolog/log"
)

func Run(cfg config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	repo, err := memory.New(domain.QuoteList, rand.Intn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed create repo")
	}

	service := quote.NewService(cfg, repo)

	tcpServer := tcp.New(cfg, service, pow.NewHashCashProvider())

	go tcpServer.Run(ctx)

	<-ctx.Done()
	log.Info().Msg("app received interrupt signal")
}
