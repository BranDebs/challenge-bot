package bot

import (
	"context"
	"log"
	"strings"

	"github.com/BranDebs/challenge-bot/ui"

	"github.com/BranDebs/challenge-bot/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot  *tgbotapi.BotAPI
	repo repository.Repository
}

func New(apiToken string, repo repository.Repository) *Bot {
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Panic(err)
		return nil
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	bot.Debug = true
	return &Bot{
		bot:  bot,
		repo: repo,
	}
}

func (b *Bot) Listen() error {
	log.Printf("Bot is listening for messages.")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		log.Printf("Unable to get updates.")
		return err
	}

	messageHandler := initMessageHandler()

	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Text {
			case ui.StartBot:
				msg = messageHandler.GetMainScreenMsg(msg)
			case ui.JoinAChallenge:
				msg = messageHandler.GetAvailableChallengesMsg(msg)
			case ui.ViewYourChallenges:
				msg = messageHandler.GetUserChallengesMsg(msg, update.Message.From.ID)
			}

			b.bot.Send(msg)

		} else if update.CallbackQuery != nil {
			text := update.CallbackQuery.Message.Text
			if strings.Contains(text, "Available Challenges") {
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
				msg = messageHandler.JoinChallengeIdMsg(msg, update.CallbackQuery)
				b.bot.Send(msg)

			} else if strings.Contains(text, "Your Challenges") {
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "")
				msg = messageHandler.ShowChallengeIdMsg(msg, update.CallbackQuery)
				b.bot.Send(msg)
			}

		}

	}

	return nil
}

func initMessageHandler() ui.Message {
	ctx := context.Background()
	keyboardProvider := ui.NewKeyboardProvider()
	textInfoProvider := ui.NewTextInfoProvider()
	return ui.NewMessage(ctx, keyboardProvider, textInfoProvider, nil)
}
