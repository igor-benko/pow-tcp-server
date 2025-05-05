package tcp

import (
	"context"
	"fmt"
	"net"

	"github.com/igor-benko/pow-tcp-server/internal/config"
	"github.com/igor-benko/pow-tcp-server/internal/service"
	"github.com/igor-benko/pow-tcp-server/pkg/pow"
	"github.com/rs/zerolog/log"
)

type Server struct {
	cfg         config.Config
	quote       service.Quote
	powProvider pow.Provider
}

func New(cfg config.Config, quote service.Quote, powProvider pow.Provider) *Server {
	return &Server{
		cfg:         cfg,
		quote:       quote,
		powProvider: powProvider,
	}
}

func (s *Server) Run(ctx context.Context) {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.Server.Port))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to start TCP server")
	}

	defer listen.Close()

	log.Info().Int("port", s.cfg.Server.Port).Msg("Server started")

	go func() {
		<-ctx.Done()
		log.Info().Msg("Shutting down server...")
		listen.Close()
	}()

	for {
		conn, err := listen.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				log.Info().Msg("Graceful shutdown complete")
				return
			default:
				log.Error().Err(err).Msg("Failed to accept connection")
			}
			continue
		}

		wrapped := POW(s.handler, s.powProvider, s.cfg.Pow.ChallengeDifficulty)
		wrapped = Metrics(wrapped)
		wrapped = Logging(wrapped)

		go wrapped(ctx, conn)
	}
}

func (s *Server) handler(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	quote, err := s.quote.GetRandomQuote(ctx)
	if err != nil {
		fmt.Fprintln(conn, err)
		return
	}

	if _, err := fmt.Fprintf(conn, "OK %s\n", quote.Content); err != nil {
		log.Error().Err(err).Msg("failed to send quote")
		return
	}
}
