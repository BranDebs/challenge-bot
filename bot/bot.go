package bot

import (
	"context"
	"log"

	"github.com/BranDebs/challenge-bot/command"
	"github.com/BranDebs/challenge-bot/command/base"

	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/logic"

	"github.com/BranDebs/challenge-bot/ui"

	"github.com/BranDebs/challenge-bot/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot  *tgbotapi.BotAPI
	repo repository.Repository
}

const (
	parseMode = "MarkdownV2"
)

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
	updates.Clear()

	ctx := context.Background()

	l := logic.New(b.repo)
	v := validator.NewValidator()
	cmdFactory := command.NewFactory()

	for update := range updates {
		if update.Message != nil {
			log.Printf("%+v", *update.Message)
			replyMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			replyMsg.ParseMode = parseMode

			cmd, err := cmdFactory.GetCommand(
				base.MsgData{
					Msg:    update.Message.Text,
					UserID: uint64(update.Message.From.ID),
				},
				l, v)
			if err != nil {
				replyMsg.Text = err.Error()
				_, _ = b.bot.Send(replyMsg)
				log.Printf("Unable to get command from factory err: %v", err)
				continue
			}

			replyMsgString, err := cmd.Execute(ctx)
			if err != nil {
				replyMsg.Text = err.Error()
				_, _ = b.bot.Send(replyMsg)
				log.Printf("Unable texecute command err: %v", err)
				continue
			}

			replyMsg.Text = base.CleanMarkdownMsg(replyMsgString)
			_, err = b.bot.Send(replyMsg)
			if err != nil {
				log.Printf("Failed to send reply to bot client err: %v")
			}
		}
	}
	return nil
}

func (b *Bot) initMessageHandler() ui.Message {
	handler := logic.New(b.repo)
	ctx := context.Background()
	keyboardProvider := ui.NewKeyboardProvider()
	textInfoProvider := ui.NewTextInfoProvider()
	return ui.NewMessage(ctx, keyboardProvider, textInfoProvider, handler)
}
