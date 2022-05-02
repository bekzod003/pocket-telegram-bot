package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const commandStart = "start"

// Handle message from  telegram bot
func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case commandStart:
		println("Start command")
		msg := tgbotapi.NewMessage(message.Chat.ID, "Hello, "+message.From.UserName)
		b.bot.Send(msg)
	}
}

// Handle message from  telegram bot
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	b.bot.Send(msg)
}
