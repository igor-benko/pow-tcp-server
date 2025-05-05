package storage

import (
	"context"

	"github.com/igor-benko/pow-tcp-server/internal/domain"
)

type Storage interface {
	Quote
}

type Quote interface {
	GetRandomQuote(ctx context.Context) (*domain.Quote, error)
}
