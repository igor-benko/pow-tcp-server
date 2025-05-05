package pow

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type HashCashProvider struct {
}

func NewHashCashProvider() *HashCashProvider {
	return &HashCashProvider{}
}

func (p *HashCashProvider) GenerateChallenge() (string, error) {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}

func (p *HashCashProvider) Validate(challenge string, difficulty int, nonce string) bool {
	prefix := strings.Repeat("0", difficulty)
	candidate := fmt.Sprintf("%s:%s", challenge, nonce)
	hash := sha256.Sum256([]byte(candidate))
	hexHash := hex.EncodeToString(hash[:])
	return strings.HasPrefix(hexHash, prefix)
}

func (p *HashCashProvider) Solve(challenge string, difficulty int) string {
	prefix := strings.Repeat("0", difficulty)
	for i := 0; ; i++ {
		candidate := fmt.Sprintf("%s:%d", challenge, i)
		hash := sha256.Sum256([]byte(candidate))
		if strings.HasPrefix(hex.EncodeToString(hash[:]), prefix) {
			return fmt.Sprintf("%d", i)
		}
	}
}
