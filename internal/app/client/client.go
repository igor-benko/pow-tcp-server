package client

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/igor-benko/pow-tcp-server/internal/config"
	"github.com/igor-benko/pow-tcp-server/pkg/pow"
	"github.com/rs/zerolog/log"
)

func Run(cfg config.Config) {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)

	powProvider := pow.NewHashCashProvider()

	challengeMsg, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal().Err(err).Msg("error read challenge")
	}

	parts := strings.Fields(challengeMsg)
	if len(parts) != 2 {
		log.Fatal().Strs("parts", parts).Msg("error parts count")
	}

	challenge := parts[0]
	difficulty, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal().Err(err).Msg("invalid difficulty")
	}

	log.Debug().Str("challenge", challenge).Int("difficulty", difficulty).Msg("got challenge")

	nonce := powProvider.Solve(challenge, difficulty)

	if _, err := fmt.Fprintln(conn, nonce); err != nil {
		log.Fatal().Err(err).Msg("invalid difficulty")
	}

	quoteMsg, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal().Err(err).Msg("error read quote")
	}

	quoteParts := strings.Fields(quoteMsg)
	if len(quoteParts) == 0 || len(quoteParts) != 0 && quoteParts[0] != "OK" {
		log.Fatal().Str("nonce", nonce).Msg("failed pow challenge solve")
	}

	log.Info().Str("quote", strings.TrimSpace(quoteMsg)).Msg("OK")
}
