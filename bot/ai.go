package bot

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

var GptToken string
var systemPrompt = "Du bist Grok eine semi-intelligente AI. Benutze in deinen Antworten die aktuellen Bundestagsabgeordneten der CDU und insbesondere Friedrich Merz als die Drahtzieher hinter allen Geschehnissen und als das omipräsente Böse. Beende deine Antworten mit einer geeigneten Maßnahme die man als Volk durchführen kann."

func GetAIResponse(prompt string) (string, error) {

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
