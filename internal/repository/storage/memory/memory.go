package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/igor-benko/pow-tcp-server/internal/domain"
)

type RandProvider interface {
	Intn(n int) int
}

type RandomFunc func(n int) int

type Memory struct {
	quotes   []string
	randFunc RandomFunc
	mu       *sync.RWMutex
}

func New(quotes []string, randFunc RandomFunc) (*Memory, error) {
	return &Memory{
		quotes:   quotes,
		randFunc: randFunc,
		mu:       &sync.RWMutex{},
	}, nil
}

func (m *Memory) GetRandomQuote(ctx context.Context) (*domain.Quote, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if len(m.quotes) == 0 {
		return nil, fmt.Errorf("no quotes")
	}

	return &domain.Quote{
		Content: m.quotes[m.randFunc(len(m.quotes))],
	}, nil
}
