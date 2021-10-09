package main

import (
	"log"

	"github.com/BranDebs/challenge-bot/bot"
	"github.com/BranDebs/challenge-bot/secrets"
)

func main() {
	apiToken, err := secrets.APIToken("apitoken")
	if err != nil {
		log.Panic(err)
	}

	botClient := bot.New(apiToken)
	err = botClient.Listen()
	if err != nil {
		log.Fatalf("Bot stopped listening err: %v", err)
	}
}
