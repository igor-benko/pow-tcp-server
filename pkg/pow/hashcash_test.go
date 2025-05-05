package pow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate_ValidSolution(t *testing.T) {
	provider := NewHashCashProvider()

	challenge := "295f73ec24f20bb9e5e5c2b6c5f6e842"
	difficulty := 3
	validNonce := "7050"

	assert.True(t, provider.Validate(challenge, difficulty, validNonce), "expected valid PoW solution to pass validation")
}

func TestValidate_InvalidSolution(t *testing.T) {
	provider := NewHashCashProvider()

	challenge := "abc123"
	difficulty := 4
	invalidNonce := "not-valid"

	assert.False(t, provider.Validate(challenge, difficulty, invalidNonce), "expected invalid PoW solution to fail validation")
}

func TestValidate_EmptyNonce(t *testing.T) {
	provider := NewHashCashProvider()

	challenge := "abc123"
	difficulty := 2
	nonce := ""

	assert.False(t, provider.Validate(challenge, difficulty, nonce), "empty nonce should not be valid")
}
