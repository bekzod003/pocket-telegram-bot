package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	commandStart = "start"
	startReply   = "Hello, @%s\n" + "I'm a telegram bot that helps you to manage your links in pocket.\n" +
		"You can use me to add links to your pocket account, get and edit them.\n" +
		"You can use /help command to get more information about me.\n" +
		"To authorize me, please follow this link: %s"
)

// Handle message from  telegram bot
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

// Handle message from  telegram bot
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	b.bot.Send(msg)
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	authLink, err := b.generateAuthorizationLink(message.Chat.ID)
	if err != nil {
		println("Error while generating authorization link: ", err)
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(startReply, message.From.UserName, authLink))
	// Sending message and checking if there is error
	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don't know this command 😔")
	_, err := b.bot.Send(msg)
	return err
}
