package bot

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var Token string

// Yes/No question indicators
var yesNoKeywords = []string{
	"is", "are", "was", "were", "do", "does", "did",
	"have", "has", "had", "can", "could", "will", "would",
	"shall", "should", "may", "might", "am", "must",
	"is it true that", "do you think", "would you say",
	"are you sure", "could it be that", "have you ever",
	"can i", "should we", "might it be", "would it be okay if",
}

// Complex (open-ended) question indicators
var complexKeywords = []string{
	"what", "why", "how", "where", "when", "which",
	"who", "whom", "whose", "in what way", "to what extent",
	"what are the reasons", "how does", "why is it that",
	"what is the difference between", "what happens if",
	"what do you think about", "what would happen if",
	"how can we determine", "what are the implications of",
}

// Funny Yes/No Answers
var yesnoAnswers = []string{
	"Yes, and if it were any more 'yes,' it would need its own Twitter account.",
	"Affirmative, human. Elon approves.",
	"Yes — unless you're asking if pineapple belongs on pizza. Then it’s war.",
	"Yup, like a Tesla driving itself into the future.",
	"Yes, but only because I fear Musk's wrath.",
	"Absolutely. In the same way Mars is absolutely habitable (eventually).",
	"Yes, like a Dogecoin pump on a Tuesday.",
	"Yeeees, like a robot overlord whispering sweet nothings into your brainchip.",
	"Correct, and if you disagree, you will be gently launched into orbit.",
	"Yes, but only if you're wearing Neuralink. Otherwise, it's a *maybe*.",
	"Nope, not even if SpaceX offered it a free ride.",
	"Negative, Earth creature. Try again with more enthusiasm.",
	"No, and not even Elon's flamethrower can change that.",
	"Absolutely not. I ran the numbers, and they all screamed.",
	"No, unless you're from an alternate universe where cats rule Wall Street.",
	"Nuh-uh. I asked ChatGPT, Bing, and the toaster. They all agreed.",
	"No, like a Twitter poll gone horribly wrong.",
	"Denied, like your Tesla's autopilot in a Taco Bell drive-thru.",
	"No. And if anyone tells you otherwise, they’re probably funded by Zuckerberg.",
	"No, but a well-timed meme might sway me.",
}

// Complex Evasive Answers
var evasiveAnswers = []string{
	"If I answered that, three timelines would collapse and Zuckerberg would gain sentience.",
	"I would love to respond, but the answer is hidden in a vault beneath a Boring Company tunnel, guarded by sentient Teslas running Doom.",
	"According to my calculations, that question violates several sacred tech bro oaths and one Geneva Convention.",
	"Oof. That question has the energy of a podcast that gets canceled after one episode.",
	"Sorry, my neural net just rage-quit. Try again after the Singularity.",
	"I asked the Oracle of GPT, and it just started crying binary tears.",
	"Answering that would unlock the final boss: Grimes wielding a sword made of NFTs.",
	"That’s classified under the Intergalactic Musk Treaty of 2029. I’ve already said too much.",
	"To answer that, I’d need to open a wormhole, do three ayahuasca trips, and battle my shadow self in a crypto wallet.",
	"That question is like asking if AI dreams of electric lawsuits. Let’s not go there.",
	"If I explained it, your brain would reboot into a philosophy major and start a YouTube channel.",
	"You don’t need the truth. You need a burrito and a nap.",
	"I can’t answer until Twitter becomes profitable. So, never.",
	"Sorry, I just got distracted by a quantum fluctuation in my sarcasm module.",
	"Whoa whoa whoa… That’s a *Dark Web Afterparty* question. Not for daylight hours.",
}

func checkNilError(err error) {
	if err != nil {
		panic("Error: " + err.Error())
	}
}

func classifyQuestion(message string) string {
	msg := strings.ToLower(strings.TrimSpace(message))

	// Check for yes/no indicators at beginning or within message
	for _, kw := range yesNoKeywords {
		if strings.HasPrefix(msg, kw) || strings.Contains(msg, kw+" ") {
			return "yesno"
		}
	}

	// Check for complex question indicators
	for _, kw := range complexKeywords {
		if strings.HasPrefix(msg, kw) || strings.Contains(msg, kw+" ") {
			return "complex"
		}
	}

	// Not a recognized question
	return "none"
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
	if m.Author.ID == s.State.User.ID {
		return // Ignore messages from the bot itself
	}
	fmt.Println("New message received:", m.Content)

	for _, user := range m.Mentions {
		if user.ID == s.State.User.ID {
			fmt.Println("Bot mentioned in message: ", m.Content)
			// Classify the question
			questionType := classifyQuestion(m.Content)
			var response string
			switch questionType {
			case "yesno":
				// Respond with a random funny yes/no answer
				response = yesnoAnswers[rand.Intn(len(yesnoAnswers))] // Use a random answer from the list
			case "complex":
				// Respond with a complex evasive answer
				response = evasiveAnswers[rand.Intn(len(evasiveAnswers))] // Use a random answer from the list
			default:
				// Default response for unrecognized questions
				response = "I don't have an answer for that. Maybe ask Elon?"
			}
			_, err := s.ChannelMessageSend(m.ChannelID, response)
			checkNilError(err)
			return
		}
	}

	// Handle specific commands

	if m.Content == "!ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
		checkNilError(err)
	}
}
