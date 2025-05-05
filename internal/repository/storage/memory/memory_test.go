package memory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomQuote(t *testing.T) {
	quotes := []string{"Life is beautiful", "The only limit is your mind", "Stay positive"}
	randFunc := func(n int) int {
		return 1
	}

	storage, err := New(quotes, randFunc)
	assert.NoError(t, err)

	ctx := context.Background()
	quote, err := storage.GetRandomQuote(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, quote)
	assert.Equal(t, "The only limit is your mind", quote.Content)
}
