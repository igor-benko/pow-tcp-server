package pow

type Provider interface {
	GenerateChallenge() (string, error)
	Validate(challenge string, difficulty int, nonce string) bool
	Solve(challenge string, difficulty int) string
}
