package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

type MessageCategory string

const (
	Uninterrested           MessageCategory = "Uninterrested"
	Mentioned               MessageCategory = "Mentioned"
	Reply                   MessageCategory = "Reply"
	ReplyWithAttachment     MessageCategory = "ReplyWithAttachment"
	MentionedWithAttachment MessageCategory = "MentionedWithAttachment"
	MentionedWithLink       MessageCategory = "MentionedWithLink"
	ReplyWithLink           MessageCategory = "ReplyWithLink"
)

func catogorizeMessage(message *discordgo.MessageCreate, session *discordgo.Session) MessageCategory {
	// Ignore messages from the bot itself
	if message.Author.ID == session.State.User.ID {
		return Uninterrested
	}

	// If the bot is mentioned, return Mentioned
	for _, user := range message.Mentions {
		if user.ID == session.State.User.ID {
			// Check if the message has an attachment
			if len(message.Attachments) > 0 {
				// If the bot is mentioned with an attachment, return MentionedWithAttachment
				return MentionedWithAttachment
			}

			if message.MessageReference != nil {
				// If the bot is mentioned in a reply, return Mentioned
				return Reply
			}

			// If the bot is mentioned without an attachment, return Mentioned
			return Mentioned
		}
	}

	// If the bot is not mentioned and not a reply, return uninterested
	return Uninterrested
}

func detectLink(message *discordgo.MessageCreate) bool {
	// Check if the message contains a link
	for _, attachment := range message.Attachments {
		if attachment.URL != "" {
			return true
		}
	}
	if strings.Contains(message.Content, "http://") || strings.Contains(message.Content, "https://") {
		return true
	}
	return false
}
