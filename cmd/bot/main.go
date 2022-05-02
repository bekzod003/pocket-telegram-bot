package main

import (
	"log"

	"github.com/bekzod003/pocket-telegram-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5312336862:AAFFtOowzF1y0gCEa5jUmxxfYMfIA21kZ10")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	pocketClient, err := pocket.NewClient("101889-bd7bf15b13d693a7d6247d4")
	if err != nil {
		log.Fatal("Error while creating pocket client: ", err)
	}

	if err := telegram.NewBot(bot, pocketClient, "localhost").Start(); err != nil {
		log.Fatal("Error while starting bot: ", err)
	}
}
