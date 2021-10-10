package main

import (
	"flag"
	"log"

	"github.com/BranDebs/challenge-bot/bot"
	"github.com/BranDebs/challenge-bot/cli"
	"github.com/BranDebs/challenge-bot/config"
	"github.com/BranDebs/challenge-bot/logic"
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

	mode := flag.String("mode", "bot", "Selects from mode: bot, cli.")
	flag.Parse()

	switch *mode {
	case "bot":
		bot, err := setupBot(conf, repo)
		if err != nil {
			log.Panic(err)
		}

		if err := bot.Listen(); err != nil {
			log.Fatalf("Bot failed to listen err: %v", err)
			return
		}
	case "cli":
		l := logic.New(repo)
		c := cli.New(l)
		if err := c.Listen(); err != nil {
			log.Fatalf("CLI failed to listen err: %v", err)
		}
	default:
		panic("wrong mode selected")
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
