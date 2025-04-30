package ai

import (
    "context"
    "fmt"
    "strings"

    openai "github.com/sashabaranov/go-openai"
)

// Client wraps the go-openai client for easier testing and customization.
type Client struct {
    api *openai.Client
}

// NewClient creates a new AI client with the provided OpenAI API key.
func NewClient(apiKey string) *Client {
    return &Client{
        api: openai.NewClient(apiKey),
    }
}

// GenerateCommitMessage sends the staged diff to GPT-3.5 and returns a one-line commit message.
// If the diff is empty, it returns an empty string without error.
func (c *Client) GenerateCommitMessage(diff string) (string, error) {
    diff = strings.TrimSpace(diff)
    if diff == "" {
        return "", nil
    }

    req := openai.ChatCompletionRequest{
        Model: openai.GPT3Dot5Turbo,
        Messages: []openai.ChatCompletionMessage{
            {
                Role:    openai.ChatMessageRoleSystem,
                Content: "You are a helpful assistant that writes concise git commit messages.",
            },
            {
                Role:    openai.ChatMessageRoleUser,
                Content: fmt.Sprintf("Write a one-line commit message (use types like feat/fix/docs) for the following diff:\n\n%s", diff),
            },
        },
        MaxTokens:   60,
        Temperature: 0.2,
    }

    resp, err := c.api.CreateChatCompletion(context.Background(), req)
    
    if err != nil {
        return "", fmt.Errorf("OpenAI API error: %w", err)
    }

    // Take the first line of the response and trim whitespace
    content := strings.TrimSpace(resp.Choices[0].Message.Content)
    firstLine := strings.SplitN(content, "\n", 2)[0]
    return firstLine, nil
}
