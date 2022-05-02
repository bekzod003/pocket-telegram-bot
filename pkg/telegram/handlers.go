package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const commandStart = "start"

// Handle message from  telegram bot
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know this command 😔")

	switch message.Command() {
	case commandStart:
		println("Start command")
		msg.Text = "Hello, @" + message.From.UserName
		// If user has no username, we should use first name
		if message.From.UserName == "" {
			msg.Text = "Hello, " + message.From.FirstName
		}

		// Sending message and checking if there is error
		_, err := b.bot.Send(msg)
		return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

// Handle message from  telegram bot
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	b.bot.Send(msg)
}
