package quote

import (
	"context"

	"github.com/igor-benko/pow-tcp-server/internal/config"
	"github.com/igor-benko/pow-tcp-server/internal/domain"
	"github.com/igor-benko/pow-tcp-server/internal/service/adapters/storage"
)

type Service struct {
	cfg     config.Config
	storage storage.Storage
}

func (s Service) GetRandomQuote(ctx context.Context) (*domain.Quote, error) {
	return s.storage.GetRandomQuote(ctx)
}

func NewService(
	cfg config.Config,
	storage storage.Storage,
) Service {
	return Service{
		cfg:     cfg,
		storage: storage,
	}
}
