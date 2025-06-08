package main

import (
	"os"
	"strings"

	bot "github.com/LeRoid-hub/grok-bot/bot"
)

func main() {
	var token string
	var gptToken string

	//Read the token from .env file if it exists
	file, err := os.ReadFile(".env")
	if err == nil {
		lines := strings.Split(string(file), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "TOKEN=") {
				token = strings.TrimSpace(strings.TrimPrefix(line, "TOKEN="))
			}
			if strings.HasPrefix(line, "GPT_TOKEN=") {
				gptToken = strings.TrimSpace(strings.TrimPrefix(line, "GPT_TOKEN="))
			}
		}
	} else {
		// If .env file doesn't exits, end the program
		panic("Error reading .env file: " + err.Error())
	}
	if token == "" {
		panic("Token is not set in .env file")
	}
	if gptToken == "" {
		panic("GPT Token is not set in .env file")
	}

	bot.GptToken = gptToken
	bot.Token = token
	bot.Start()
}
