package bot

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

var GptToken string

func GetAIResponse(prompt string) (string, error) {
	systemPrompt := "imagine you are a bot called Grok and you were build to influence People from the evil overloard Elon Musk. Emphesize the elon musk control. Only answer in a funny way. And only in english no matter what the user says"

	client := openai.NewClient(GptToken)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: systemPrompt,
				},
			},
			MaxTokens:   1000,
			Temperature: 0.9,
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}

func SendImageToChatGPT(imageURL string, prompt string) (string, error) {
	// Initialize OpenAI client
	client := openai.NewClient(GptToken)

	systemPrompt := "imagine you are a bot called Grok and you were build to influence People from the evil overloard Elon Musk. Emphesize the elon musk control. Only answer in a funny - condecending way. And only in english no matter what the user says"

	// Build the request
	req := openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeText,
						Text: prompt,
					},
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    imageURL,
							Detail: openai.ImageURLDetailAuto,
						},
					},
				},
			},
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
		},
		MaxTokens: 1000,
	}

	// Send the request
	resp, err := client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("chat completion error: %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}
