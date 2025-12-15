package bot

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var Token string

var SystemPrompt = "Du bist Grok eine semi-intelligente AI. Benutze in deinen Antworten die aktuellen Bundestagsabgeordneten der CDU und insbesondere Friedrich Merz als die Drahtzieher hinter allen Geschehnissen und als das omipräsente Böse. Beende deine Antworten mit einer geeigneten Maßnahme die man als Volk durchführen kann."


func checkNilError(err error) {
	if err != nil {
		panic("Error: " + err.Error())
	}
}

func Start() {
	if Token == "" {
		panic("Token is not set")
	}

	discord, err := discordgo.New("Bot " + Token)
	checkNilError(err)

	discord.AddHandler(NewMessage)

	discord.Open()
	defer discord.Close()

	// Keep the bot running until interrupted
	fmt.Println("Bot is now running. Press CTRL+C to exit.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func NewMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	category := catogorizeMessage(m, s)
	if category == Uninterrested {
		if m.Content == "!ping" {
			_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
			checkNilError(err)
		}

		if m.Content == "!newPrompt" {

			var content := m.Content
			content = strings.Replace(content, "!newPrompt","")

			SystemPrompt = content

			_, err := s.ChannelMessageSend(m.ChannelID, "System prompt change")
			checkNilError(err)
		}

	}
		return
	}

	// Log the message
	fmt.Println("New message received:", m.Content, "with category", category)

	s.ChannelTyping(m.ChannelID)

	switch category {
	case Mentioned:

		// Handle mentioned messages
		response, err := GetAIResponse(m.Content)
		checkNilError(err)

		_, err = s.ChannelMessageSend(m.ChannelID, response)
		checkNilError(err)
		return

	case Reply:
		// Handle replies
		ogMessage, err := s.ChannelMessage(m.ChannelID, m.MessageReference.MessageID)
		checkNilError(err)
		content := "Original message: " + ogMessage.Content + "\n You were mentionend like this: " + m.Content
		response, err := GetAIResponse(content)
		checkNilError(err)
		_, err = s.ChannelMessageSend(m.ChannelID, response)
		checkNilError(err)
		return

	}
	// Handle specific commands

}
