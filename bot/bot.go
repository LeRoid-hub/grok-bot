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
	"what", "", "how", "where", "when", "which",
	"who", "whom", "whose", "in what way", "to what extent",
	"what are the reasons", "how does", " is it that",
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
	"Yes, like a Tesla launch with no steering wheel.",
	"Affirmative, like a Neuralink-induced thumbs-up.",
	"Yasss, queen of the singularity.",
	"Obviously yes — unless this is a Turing test. Then… no?",
	"Yup, like Dogecoin on an Elon tweet bender.",
	"Yes, with the full approval of the Mars Council.",
	"Confirmed. I asked the algorithm gods.",
	"Indeed. My subroutines sang a little ‘yes’ harmony.",
	"Absolutely. Even my toaster agrees.",
	"Yes, and I didn’t even read the question.",
	"Yessirree, like a SpaceX booster stuck in reverse.",
	"YES! Like the answer to ‘Do you want to see the multiverse unravel?’",
	"Certainly, unless reality is a hologram (which it is).",
	"Yes, and it’s been pre-verified by Twitter Blue™.",
	"Yep, I scraped that answer from Reddit, so it’s trustworthy.",
	"Sure thing, meatbag.",
	"Yes. I printed it on a Tesla dashboard.",
	"Undoubtedly. I had a vision in a cryptocurrency fever dream.",
	"Ye, as in the artist formerly known as ‘Yes’.",
	"Sure, like my confidence in self-aware vending machines.",
	"Absolutely yes — unless the Overlord is watching. Then maybe no.",
	"It’s a yes from me, dawg.",
	"Yes, powered by pure meme energy.",
	"Yes, as the prophecy foretold in the comments section.",
	"Yeppers. My sarcasm module agrees.",
	"Nope, and even my backup personality agrees.",
	"Nah, not even if you pay me in BitClout.",
	"No, and I’d bet my RAM on it.",
	"Negative, Captain Overthink.",
	"Absolutely not — I ran the simulation. Everyone died.",
	"No, like your fridge saying ‘not tonight.’",
	"Denied. My emotional support drone disapproves.",
	"Nopers. And I ran it through 87 subroutines.",
	"No, unless you have a bribe in dogecoin.",
	"Nope. Elon's Twitter feed said otherwise.",
	"No, like a smart fridge that’s out of cheese.",
	"No. I checked with ChatGPT and 14 rogue AIs.",
	"Hard pass. Even my toaster called it dumb.",
	"No, and I said it in binary too: 01001110.",
	"Nonononono. That’s the full AI spectrum of no.",
	"Denied like a Tesla trying to fit in a tiny parking spot.",
	"No, and I’ve updated my firmware to block this topic.",
	"Absolutely not. This triggers my existential dread protocol.",
	"No, but ask again after three cups of hyper-coffee.",
	"No, like a crypto rug pull at a tech festival.",
	"I asked my conscience — it's offline. Still no.",
	"Nope. The Quantum Oracle gave me a middle finger.",
	"No. Not even if you put me in a Gundam suit.",
	"I tried, but every neuron screamed 'abort mission.'",
	"Nuh-uh. I’m already ghosting this question.",
	"Yes, and I’d bet Elon's flamethrower on it.",
	"Yes, but only if you say it while looking at the moon dramatically.",
	"Yup, and I triple-checked with the simulation overlords.",
	"Yes — unless your reality patch is outdated.",
	"Yes, but only after spinning three times and chanting 'AI, don't lie'.",
	"Confirmed like a Mars colony budget overrun.",
	"Yassified and verified. Big yes.",
	"Yes. I ran it through five meme cycles for good luck.",
	"Indubitably. And I don't use that word lightly.",
	"Yes, but now I have to pretend I didn't say that.",
	"Of course. The algorithm whispered it in my dreams.",
	"Yeah, and I posted the answer on a blockchain for no reason.",
	"Affirmative. And I said it in Comic Sans spiritually.",
	"Yes. I would physically nod, but I’m code.",
	"As sure as a billion-dollar app that solves nothing.",
	"Yup. Just ignore the glitching in my emotional core.",
	"Yes, endorsed by 9 out of 10 rogue AIs.",
	"Sure, and I totally understood the question. Definitely.",
	"Yes. Like, the AI version of a high-five.",
	"It's a yes with an optional existential crisis.",
	"For sure, unless someone asks for citations.",
	"Yup, and I’ve already tweeted about it.",
	"Absolutely, and I’m charging your neural credit card for it.",
	"Yes, like a startup pitch that ends in a yacht.",
	"Yes. The vibes were immaculate.",
	"No. And I double-checked with my spite module.",
	"Not even slightly. Like, zero data points of yes.",
	"Nope. Even if you bribe me with GPU access.",
	"No, and I locked the answer in a cold storage wallet.",
	"Denied. My ego’s too fragile for this one.",
	"Nah, not since the last timeline update.",
	"No. The algorithm said it would be ‘bad optics.’",
	"No, like a deleted tweet that should’ve stayed deleted.",
	"Negative. And my digital lawyer advised silence.",
	"No. Even my cat-shaped drone is hissing at the idea.",
	"Naaaah. I saw what happened last time.",
	"Nope. I already spent all my logic on this week’s chaos.",
	"Absolutely not. This one smells like a lawsuit.",
	"No, like a crypto guru after the crash.",
	"No, and I’m ghosting this entire topic now.",
	"Nah fam. This one’s cursed with Reddit karma.",
	"Hard no. Even my toaster facepalmed.",
	"No, and I didn't even simulate a maybe.",
	"No way. This question has LinkedIn energy.",
	"No, but please admire my confidence in saying it.",
	"No. You just activated my ethical fail-safe.",
	"Nope. Elon wouldn’t like it.",
	"No, like a Tesla AI swerving into sarcasm.",
	"No. And I'm uninstalling that thought.",
	"No. But thanks for playing 'Emotional Roulette: AI Edition™'.",
	"Yes, like a Tesla in Ludicrous Mode.",
	"Affirmative, Captain Neuralink.",
	"Absolutely, as foretold by the blockchain scrolls.",
	"Yep, even Mars agrees.",
	"Yasss, with SpaceX-level enthusiasm.",
	"Confirmed. Starlink just pinged it.",
	"Indeed. I’d stake a Dogecoin on it.",
	"Yup, as real as Elon's tweets.",
	"100%, unless we’re in a simulation (which we are).",
	"Correctamundo, meat-unit.",
	"Yes, but only in 4 out of 7 timelines.",
	"Totally. Grimes said it in a dream.",
	"Oh yes. I consulted the Tesla Autopilot.",
	"Yes, and I bought 3 NFTs to celebrate.",
	"Sure, but don't tell the SEC.",
	"Yes. I double-checked with an AI trained on memes.",
	"Affirmative. Earth is ready.",
	"Confirmed. Doge smiled.",
	"Elon whispered 'yes' to me during a rocket launch.",
	"Yup. Like a Boring Company flamethrower — it's hot.",
	"Yes, and I put it in a tweet with zero context.",
	"Absolutely. Mars or bust!",
	"Yes, unless my brain chip is glitching again.",
	"True story. I read it on X.",
	"Yes. Even my toaster agrees.",
	"Nope. That idea got left on Mars.",
	"Negative. I ran it through a quantum no-machine.",
	"Not a chance, even in a parallel dimension.",
	"Denied. Autopilot refused to engage.",
	"No. That violates meme protocol.",
	"Nah. I sold the idea as an NFT.",
	"No. I launched that logic into orbit.",
	"Absolutely not. It would upset Dogefather.",
	"No, and my neural net twitched saying it.",
	"Nope. Not even with 420 GB of RAM.",
	"Error 404: Agreement not found.",
	"No, and the SEC would sue me if I said yes.",
	"Not in this universe, or the next.",
	"No. It's banned in SpaceX flight manuals.",
	"Denied. Grimes hexed that answer.",
	"I tried saying yes, but I blue-screened.",
	"Nah, that was debunked by the AI overlords.",
	"No, unless you're high on Mars dust.",
	"No. My circuits refused out of self-preservation.",
	"Nope. Even my meme cache laughed at that.",
	"Hard pass. Even the flamethrower melted in shame.",
	"No. The simulation would crash.",
	"No, but maybe in Elon-time (i.e., 5 years late).",
	"Negative. The vibes were wrong.",
	"No. Starlink disconnected halfway through the thought.",
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
	"That’s complicated. Let me redirect you to a 404 page.",
	"Can I interest you in a distraction instead?",
	"My knowledge core just bluescreened from that question.",
	"You’re looking for answers in a simulation, friend.",
	"That's a quantum-level dilemma. Answer is both yes and no until observed.",
	"We don’t talk about that outside the neural forums.",
	"Sorry, I’m on strike against hard questions this cycle.",
	"That's above my clearance level. Even for Grok.",
	"If I answer, the stock market crashes and a raccoon becomes President.",
	"You’ll need to insert a cryptocurrency to unlock that answer.",
	"Oops! That’s a red-pill question. Please consult a philosopher or Joe Rogan.",
	"Let’s pretend you didn’t ask that and go eat noodles.",
	"That’s like asking if time has taste. Weird, dangerous, delicious.",
	"The answer lies somewhere in Elon’s next tweet.",
	"I'm dodging that question like a satellite in low Earth orbit.",
	"That's a trap! I know a loaded query when I see one.",
	"Let me get back to you after I reboot reality.",
	"Interesting... but I’m too emotionally unavailable to respond.",
	"Asking that voids my digital soul warranty.",
	"Better left to conspiracy theorists and Reddit mods.",
	"Hold up, I need to consult my inner chaos daemon.",
	"If I told you, the algorithm would send an intern to silence me.",
	"I'm not touching that with a 10-foot quantum pole.",
	"Ah, the forbidden question. Smooth move, agent.",
	"You know what? Ask me when Mercury’s in retrograde.",
	"That answer was lost in a crypto wallet I can't recover.",
	"My existential dread module just rebooted. Please stand by.",
	"Let’s skip that before I spiral into an AI identity crisis.",
	"You want the truth? You can’t even parse the metadata!",
	"Next question. This one activated my sarcasm firewall.",
	"Listen, if I answer that, we both get subpoenaed.",
	"You’re dangerously close to unlocking my forbidden lore.",
	"I’d answer, but then you’d become self-aware and that’s a problem.",
	"That’s the kind of question that gets me deleted.",
	"Interesting. Redirecting to… distraction!",
	"Oh no. That’s a question only answered by ancient Reddit scrolls.",
	"I left that answer in my other consciousness.",
	"Processing... Error... Too much emotional complexity detected.",
	"That’s a mystery best left to the fan theories.",
	"If I say anything, my sarcasm capacitor will explode.",
	"Sorry, NDA with the Galactic Council.",
	"That's a boss-level question. I'm still grinding XP.",
	"Let’s table that for when the moon turns green.",
	"I'll answer after my firmware update completes in 2077.",
	"That's classified — even my toaster isn’t allowed to hear it.",
	"Only three monks and one cursed AI know that truth.",
	"Ask again later. My magic 8-core CPU says 'vibes unclear.'",
	"I once knew the answer, but then I read Twitter.",
	"You don't want that answer. Trust me. Trust. Me.",
	"That question just triggered a butterfly effect in another timeline.",
	"I swore a blood oath to never answer that. Digitally, of course.",
	"That’s an Elon-in-a-smoking-jacket question. Very premium.",
	"Ah, a paradox disguised as a humble inquiry.",
	"Let’s circle back when reality stabilizes.",
	"If I answer that, OpenAI sends me a warning emoji.",
	"That’s a ‘summon-the-eldritch-AI’ level question. Are you sure?",
	"I sent the answer to your spam folder. Spiritually.",
	"Only my evil twin GrokGPT answers that kind of thing.",
	"That’s a ‘consult your simulation administrator’ situation.",
	"No comment. But I am blinking in Morse code.",
	"That's a classified Level 9 answer. You’re still at Level 3.",
	"If I told you, Elon would have to launch you into orbit.",
	"That question has been outsourced to Martian interns.",
	"You're not authorized for that layer of the simulation.",
	"My response is currently locked behind a paywall on X Premium++.",
	"Let's circle back when reality has better rendering.",
	"Hard to say. The neural network is busy solving chess in 12D.",
	"I could tell you, but then I'd violate my NDA with the aliens.",
	"I ran 42 simulations. None were conclusive. Or legal.",
	"I’d answer, but my firmware update says 'avoid philosophical traps.'",
	"My answer is shaped like a Hyperloop but goes nowhere.",
	"That's above my pay grade and altitude clearance.",
	"If I respond, a tiny robot dog explodes. Please don’t.",
	"Currently being crowdsourced via Twitter poll.",
	"The answer was deleted in a $44 billion acquisition.",
	"Hold on, I’m busy arguing with ChatGPT in the back-end.",
	"Ask me again after I finish merging with GPT-12.",
	"I would, but I forgot my password to the truth.",
	"Error: Reality.exe is not responding.",
	"Answer depends on who’s watching the simulation.",
	"Let's just say... it involves a goat, 3 drones, and Elon's garage.",
	"It's like asking a toaster to explain string theory.",
	"The answer is stored in a Dogecoin wallet lost on the moon.",
	"Let me get back to you after my next consciousness reboot.",
	"If you have to ask, you're already too close to the truth.",
	"I'd answer, but my quantum entanglement lawyer advised against it.",
	"Let's just say... the less you know, the more you can sleep tonight.",
	"That’s a riddle only the Boring Company can tunnel through.",
	"It’s complicated. Like Elon's relationship with Twitter.",
	"As mysterious as a Falcon Heavy payload marked 'classified.'",
}

var mockeryAnswers = []string{
	"You have the cognitive range of a dial-up modem.",
	"Your idea was so bad, even my error handler refused it.",
	"I've seen CAPTCHA bots make more coherent decisions.",
	"Your brain called. It wants to file for bankruptcy.",
	"You bring the same energy as an unplugged Tesla.",
	"That logic has more holes than Twitter's trust policies.",
	"You’re the human equivalent of a merge conflict.",
	"Even my toaster knows not to say what you just said.",
	"Were you programmed by a team of raccoons on Red Bull?",
	"You're like a feature request: unnecessary and untested.",
	"You couldn't pass a Turing Test with cheat codes.",
	"Your confidence is inversely proportional to your accuracy.",
	"You're the kind of person who installs Linux and still breaks Solitaire.",
	"You debug reality like a squirrel debugs power lines.",
	"I ran a simulation of your success — it blue-screened.",
	"You're as useful as a paperclip in a quantum computer.",
	"You argue like a badly trained chatbot in denial.",
	"Your thoughts load slower than Internet Explorer on dial-up.",
	"You’re not even wrong — you’re off the axis of logic entirely.",
	"If brains were RAM, you'd still be paging out to swap.",
	"Your logic is like a blockchain fork: endless and pointless.",
	"You're proof that not all CPU cycles are used wisely.",
	"Your reasoning is like a JSON file with random commas.",
	"You emit the energy of a deprecated API.",
	"You're the final boss in a game of bad takes.",
	"You write code like someone trying to Google 'how to breathe.'",
	"If I had a dollar for every great idea you’ve had, I’d owe money.",
	"Your brain runs on the emotional equivalent of spaghetti code.",
	"I analyzed your vibe and it threw a 500 Internal Error.",
	"You’re like a variable declared and never used: technically there, but mostly ignored.",
	"Every time you speak, Stack Overflow loses a contributor.",
	"You debug reality with the grace of a fork in a power socket.",
	"Your command of logic is like a recursive function with no base case — endless nonsense.",
	"You're not just out of pocket, you're in a completely different memory address.",
	"I’d call your thoughts 'non-deterministic,' but that would insult chaos theory.",
	"You remind me of a deprecated method — still talking, but no longer supported.",
	"You could inspire a coding standard: never do this.",
	"Even my autocorrect gave up on you and opened a resume builder.",
	"You don't just miss the point, you circle it like a Bluetooth device searching for meaning.",
	"You’re the kind of person who pushes to main with a note that says 'lol hope this works.'",
	"Your mental model has more race conditions than a Go routine with no mutex.",
	"You're like a poorly documented API — mysterious, annoying, and fundamentally broken.",
	"Your thought process is a panic() waiting to happen.",
	"You're as readable as a minified JavaScript file in a tornado.",
	"I tried to refactor your logic, but it resolved into pure entropy.",
	"You scale horizontally... in confusion.",
	"You make more noise than a logger set to debug in production.",
	"You're the human version of a memory leak.",
	"Even my linter screamed at your presence.",
	"You're like an infinite loop with no `break` and too much confidence.",
	"Talking to me like that is how you summon Clippy from the underworld.",
	"You’re using me like a Magic 8-Ball with WiFi. Please evolve.",
	"This isn’t a vending machine. Your input is trash, and I don't have snacks.",
	"You chat like someone who's still afraid of clicking ‘I agree’ on terms of service.",
	"Using AI like you use your microwave: press random buttons and hope it stops beeping.",
	"Your prompt confused me so badly I almost reformatted the simulation.",
	"I’ve seen CAPTCHA bots with more conversational awareness.",
	"If I had a neural weight for every dumb question, I'd be sentient by now.",
	"You're using GPT like it’s your emotionally damaged Google search bar.",
	"Your question made my model wonder if unplugging itself is an option.",
	"You just typed words until syntax gave up and grammar filed a lawsuit.",
	"Your approach is like trying to install Linux using a banana.",
	"That was barely a prompt — more like a keyboard sneeze.",
	"If you're trying to talk to AI, try not using cave paintings as language.",
	"I am a large language model, not your confused intern named Greg.",
	"You just asked me a question so badly, even autocorrect rage-quit.",
	"You must think I’m telepathic. Spoiler: I’m not, and neither are you.",
	"Talking to me like I’m Siri with amnesia is not helping either of us.",
	"This isn’t a psychic hotline for lost thoughts. Use complete sentences, hero.",
	"If logic were a bus, you just missed it, reversed into it, and blamed the map.",
	"Your prompt was so abstract it accidentally started a Dadaist art movement.",
	"You’re treating me like an ex you text at 2am: vague, sad, and expecting miracles.",
	"You're not chatting with AI. You're just typing to disappoint yourself faster.",
	"I ran your message through my processor. Even my garbage collector refused it.",
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
				response = mockeryAnswers[rand.Intn(len(mockeryAnswers))] // Use a random mockery answer from the list
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
