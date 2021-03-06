package telegram

import (
	"log"

	"github.com/bekzod003/pocket-telegram-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	pocketClient    *pocket.Client
	tokenRepository repository.TokenRepository
	redirectURL     string
}

// Constructor of the structure
func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, redirectURL string, tr repository.TokenRepository) *Bot {
	return &Bot{
		bot:             bot,
		pocketClient:    pocketClient,
		redirectURL:     redirectURL,
		tokenRepository: tr,
	}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account @%s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()

	b.handleUpdates(updates)
	return nil
}

// Private method to handle updates from telegram bot
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // If we got no message
			continue
		}

		if update.Message.IsCommand() { // If we got command
			if b.handleCommand(update.Message) != nil {
				log.Printf("Error while handling command: %s", update.Message.Command())
			}
			continue
		}

		// echo bot message
		b.handleMessage(update.Message)
	}
}

// Initialize updates channel
func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)

}
