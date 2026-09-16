package integrations_gemini

import "context"

type Generator interface {
	Generate(
		ctx context.Context,
		prompt string,
		model string,
	) ([]byte, error)
}
