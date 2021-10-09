package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot *tgbotapi.BotAPI
}

func New(apiToken string) *Bot {
	bot, err := tgbotapi.NewBotAPI(apiToken)
	if err != nil {
		log.Panic(err)
		return nil
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	bot.Debug = true
	return &Bot{
		bot: bot,
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

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		b.bot.Send(msg)
	}

	return nil
}
