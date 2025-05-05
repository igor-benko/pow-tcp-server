package service

import (
	"context"

	"github.com/igor-benko/pow-tcp-server/internal/domain"
)

type Quote interface {
	GetRandomQuote(ctx context.Context) (*domain.Quote, error)
}
