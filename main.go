package main

import (
	"os"
	"strings"

	bot "github.com/LeRoid-hub/grok-bot/bot"
)

func main() {
	var token string

	//Read the token from .env file if it exists
	file, err := os.ReadFile(".env")
	if err == nil {
		lines := strings.Split(string(file), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "TOKEN=") {
				token = strings.TrimSpace(strings.TrimPrefix(line, "TOKEN="))
				break
			}
		}
	} else {
		// If .env file doesn't exits, end the program
		panic("Error reading .env file: " + err.Error())
	}
	if token == "" {
		panic("Token is not set in .env file")
	}

	bot.Token = token
	bot.Start()
}
