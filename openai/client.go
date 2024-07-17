package openai

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

// Структура для хранения клиента OpenAI
type OpenAIClient struct {
	client *openai.Client
}

// Создание новый клиент OpenAI
func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		client: openai.NewClient(apiKey),
	}
}

// Получение ответа от OpenAI для заданного вопроса
func (c *OpenAIClient) GetResponse(ctx context.Context, question string) (string, error) {
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
	}

	resp, err := c.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
