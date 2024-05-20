package openai

import (
	"context"
	"fmt"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

type ChatGPTResponder interface {
	RequestCompletion(string, float32, string) (string, error)
}

type OpenAIClient struct {
	client *openai.Client
}

func NewClient(authToken string) *OpenAIClient {
	client := new(OpenAIClient)
	client.client = openai.NewClient(authToken)
	return client
}

// RequestCompletion, request ChatGPT for a text completion.
func (c *OpenAIClient) RequestCompletion(prompt string, temperature float32, openaiModel string) (string, error) {
	var model string
	if openaiModel == "4.0" || openaiModel == "4" || openaiModel == "4o" {
		model = openai.GPT4o
	} else {
		model = openai.GPT3Dot5Turbo
	}
	// Timeout for API call.
	timeout, err := time.ParseDuration("1m30s")
	if err != nil {
		return "", fmt.Errorf("error parsing time duration: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			// Model:       openai.GPT3Dot5Turbo,
			Model:       model,
			Temperature: temperature,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt, // prompt being sent to ChatGPT API.
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("ChatCompletion error: %w", err)
	}

	// Response aka completion received from ChatGPT.
	completion := resp.Choices[0].Message.Content

	return completion, nil
}

type MockOpenAIClient struct {
}

func NewMockClient() *MockOpenAIClient {
	return new(MockOpenAIClient)
}

func (m *MockOpenAIClient) RequestCompletion(prompt string, temperature float32, model string) (string, error) {
	return "Using Mock OpenAI API", nil
}
