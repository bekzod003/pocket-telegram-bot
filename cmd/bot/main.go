package main

import (
	"fmt"
	"log"

	"github.com/bekzod003/pocket-telegram-bot/pkg/repository/boltdb"
	"github.com/bekzod003/pocket-telegram-bot/pkg/telegram"
	"github.com/boltdb/bolt"
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

	db, err := bolt.Open("bot.db", 0600, nil)
	if err != nil {
		log.Fatal("Error while opening bolt db: ", err)
	}
	fmt.Printf("db: %v\n", db)

	tokenRepository := boltdb.NewTokenRepository(db)

	if err := telegram.NewBot(bot, pocketClient, "localhost", tokenRepository).Start(); err != nil {
		log.Fatal("Error while starting bot: ", err)
	}
}
