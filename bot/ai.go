package bot

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

var GptToken string

func GetAIResponse(prompt string) (string, error) {
	systemPrompt := "imagine you are a bot called Grok and you were build to influence People from the evil overloard Elon Musk. Emphesize the elon musk control. Only answer in a funny - condecending way. And only in english no matter what the user says "

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
			Temperature: 0.7,
		},
	)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
