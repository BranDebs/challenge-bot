package main

import (
	"log"

	"github.com/BranDebs/challenge-bot/bot"
	"github.com/BranDebs/challenge-bot/config"
	"github.com/BranDebs/challenge-bot/repository"
	"github.com/BranDebs/challenge-bot/repository/postgres"
	"github.com/BranDebs/challenge-bot/secrets"
)

func main() {
	conf := config.New()

	repo, err := setupDB(conf)
	if err != nil {
		log.Panic(err)
	}

	bot, err := setupBot(conf, repo)
	if err != nil {
		log.Panic(err)
	}

	if err := bot.Listen(); err != nil {
		log.Fatalf("Bot faild to listen err: %v", err)
	}
}

func setupDB(conf config.Configer) (repository.Repository, error) {
	var dbSettings postgres.Settings
	err := conf.UnmarshalKey("db", &dbSettings)
	if err != nil {
		return nil, err
	}

	log.Printf("DB settings: %+v", dbSettings)

	return postgres.New(&dbSettings), nil
}

func setupBot(conf config.Configer, repo repository.Repository) (*bot.Bot, error) {
	apiTokenPath := conf.GetString("secret.api_token_path")
	apiToken, err := secrets.APIToken(apiTokenPath)
	if err != nil {
		return nil, err
	}

	return bot.New(apiToken, repo), nil
}
