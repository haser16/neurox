package integrations_gemini

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/genai"
)

type Client struct {
	client *genai.Client
}

func NewClient(ctx context.Context, config *Config) (*Client, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  config.APIKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("create Gemini client: %w", err)
	}

	return &Client{
		client: client,
	}, nil
}

func (c *Client) Close() error {
	return nil
}

var _ Generator = (*Client)(nil)

var _ = time.Second
