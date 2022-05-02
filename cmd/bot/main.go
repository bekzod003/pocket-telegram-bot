package main

import (
	"log"

	"github.com/bekzod003/pocket-telegram-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5312336862:AAFFtOowzF1y0gCEa5jUmxxfYMfIA21kZ10")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	if err := telegram.NewBot(bot).Start(); err != nil {
		log.Fatal("Error while starting bot: ", err)
	}
}
