package ui

import (
	"context"
	"fmt"
	"strconv"

	"github.com/BranDebs/challenge-bot/logic"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Message interface {
	GetMainScreenMsg(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig
	GetAvailableChallengesMsg(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig
	GetUserChallengesMsg(msg tgbotapi.MessageConfig, userID int) tgbotapi.MessageConfig
	JoinChallengeIdMsg(msg tgbotapi.MessageConfig, query *tgbotapi.CallbackQuery) tgbotapi.MessageConfig
	ShowChallengeIdMsg(msg tgbotapi.MessageConfig, query *tgbotapi.CallbackQuery) tgbotapi.MessageConfig
}

type MessageImpl struct {
	ctx              context.Context
	keyboardProvider KeyboardProvider
	textInfoProvider TextInfoProvider

	challengeHandler logic.ChallengeHandler
}

func NewMessage(
	ctx context.Context,
	keyboardProvider KeyboardProvider,
	textInfoProvider TextInfoProvider,
	challengeHandler logic.ChallengeHandler,
) Message {
	return MessageImpl{
		ctx:              ctx,
		keyboardProvider: keyboardProvider,
		textInfoProvider: textInfoProvider,
		challengeHandler: challengeHandler,
	}
}

func (m MessageImpl) GetMainScreenMsg(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	msg.Text = "Get started on Challenges =)"
	msg.ReplyMarkup = m.keyboardProvider.StaticMainChallangePageKeyboard()
	return msg
}

func (m MessageImpl) GetAvailableChallengesMsg(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	// TODO: Retrieve available challenges here
	//_, err := m.challengeHandler.ListChallenges(m.ctx, nil, 0, 100)
	//if err != nil {
	//	msg.Text = "Failed to retrieve challenges"
	//	return msg
	//}

	board := m.keyboardProvider.GetChallengesKeyboard(nil)
	msg.Text = m.textInfoProvider.GetChallengesText(nil, All)
	msg.ParseMode = parseMode
	msg.ReplyMarkup = board
	return msg
}

func (m MessageImpl) GetUserChallengesMsg(msg tgbotapi.MessageConfig, userID int) tgbotapi.MessageConfig {
	// TODO: Retrieve user's enrolled challenges here
	board := m.keyboardProvider.GetChallengesKeyboard(nil)
	msg.Text = m.textInfoProvider.GetChallengesText(nil, User)
	msg.ParseMode = parseMode
	msg.ReplyMarkup = board
	return msg
}

func (m MessageImpl) JoinChallengeIdMsg(msg tgbotapi.MessageConfig, query *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	if query.Data == BackKeyword {
		return m.GetMainScreenMsg(msg)
	}

	msg.ReplyMarkup = RemoveInlineKeyboard(query.Message.Chat.ID, query.Message.MessageID)

	msg.Text = "*bold \\*text*\n_italic \\*text_\n__underline__\n~strikethrough~\n*bold _italic bold ~italic bold strikethrough~ __underline italic bold___ bold*\n[inline URL](http://www.example.com/)\n[inline mention of a user](tg://user?id=123456789)\n`inline fixed-width code`\n```\npre-formatted fixed-width code block\n```\n```python\npre-formatted fixed-width code block written in the Python programming language\n```"
	msg.Text = msg.Text + query.Data

	msg.ParseMode = parseMode
	return msg
}

func (m MessageImpl) ShowChallengeIdMsg(msg tgbotapi.MessageConfig, query *tgbotapi.CallbackQuery) tgbotapi.MessageConfig {
	if query.Data == BackKeyword {
		return m.GetMainScreenMsg(msg)
	}

	msg.ReplyMarkup = RemoveInlineKeyboard(query.Message.Chat.ID, query.Message.MessageID)

	msg.Text = fmt.Sprintf("userid=%v, challengeid=%v", query.From.ID, query.Data)
	challengeID, err := strconv.ParseUint(query.Data, 10, 64)
	if err != nil {
		msg.Text = fmt.Sprintf("Error parsing challengeID: %v, err:%v", query.Data, err)
		return msg
	}

	msg.ReplyMarkup, err = m.keyboardProvider.GetChallengeActionKeyboard(challengeID, uint64(query.From.ID))
	if err != nil {
		msg.Text = fmt.Sprintf("Error getting challenge action keyboard, err:%v", err)
		return msg
	}

	return msg
}
