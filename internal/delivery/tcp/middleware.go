package tcp

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/igor-benko/pow-tcp-server/pkg/pow"
	"github.com/rs/zerolog/log"
)

func Logging(next func(context.Context, net.Conn)) func(context.Context, net.Conn) {
	return func(ctx context.Context, conn net.Conn) {
		addr := conn.RemoteAddr().String()
		log.Info().Str("addr", addr).Msg("New connection")
		next(ctx, conn)
	}
}

func Metrics(next func(context.Context, net.Conn)) func(context.Context, net.Conn) {
	return func(ctx context.Context, conn net.Conn) {
		next(ctx, conn)
	}
}

func POW(next func(context.Context, net.Conn), provider pow.Provider, difficulty int) func(context.Context, net.Conn) {
	return func(ctx context.Context, conn net.Conn) {
		challenge, err := provider.GenerateChallenge()
		if err != nil {
			log.Error().Err(err).Msg("failed to generate challenge")
			return
		}

		reader := bufio.NewReader(conn)

		// build and send challenge
		challengeMsg := fmt.Sprintf("%s %d", challenge, difficulty)
		fmt.Fprintln(conn, challengeMsg)

		// read challenge solution
		nonce, err := reader.ReadString('\n')
		if err != nil {
			log.Error().Err(err).Msg("failed to read nonce")
			return
		}

		nonce = strings.TrimSpace(nonce)

		if !provider.Validate(challenge, difficulty, nonce) {
			fmt.Fprintln(conn, "invalid nonce")
			log.Error().Err(err).Msg("failed to solve challenge")
			return
		}

		log.Debug().Str("nonce", nonce).Msg("successfully solve challenge")

		next(ctx, conn)
	}
}
