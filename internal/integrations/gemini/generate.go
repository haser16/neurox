package integrations_gemini

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

func (c *Client) Generate(
	ctx context.Context,
	prompt string,
	model string,
) ([]byte, error) {
	result, err := c.client.Models.GenerateContent(
		ctx,
		model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("generate image: %w", err)
	}

	for _, candidate := range result.Candidates {
		if candidate.Content == nil {
			continue
		}

		for _, part := range candidate.Content.Parts {
			if part.InlineData == nil {
				continue
			}

			if len(part.InlineData.Data) == 0 {
				continue
			}

			return part.InlineData.Data, nil
		}
	}

	return nil, fmt.Errorf(
		"Gemini response does not contain generated image",
	)
}
