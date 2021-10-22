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

	logicHandler logic.Handler
}

func NewMessage(
	ctx context.Context,
	keyboardProvider KeyboardProvider,
	textInfoProvider TextInfoProvider,
	logicHandler logic.Handler,
) Message {
	return MessageImpl{
		ctx:              ctx,
		keyboardProvider: keyboardProvider,
		textInfoProvider: textInfoProvider,
		logicHandler:     logicHandler,
	}
}

func (m MessageImpl) GetMainScreenMsg(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	msg.Text = GetStartedOnChallengesMsg
	msg.ReplyMarkup = m.keyboardProvider.StaticMainChallengePageKeyboard()
	return msg
}

func (m MessageImpl) GetAvailableChallengesMsg(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	challenges, err := m.logicHandler.ListChallenges(m.ctx)
	if err != nil {
		msg.Text = "Failed to retrieve challenges"
		return msg
	}
	if len(challenges) == 0 {
		msg.Text = "No challenges available"
		return msg
	}

	board := m.keyboardProvider.GetChallengesKeyboard(challenges)
	msg.Text = m.textInfoProvider.GetChallengesText(challenges, All)
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
	_, err := strconv.ParseUint(query.Data, 10, 64)
	if err != nil {
		msg.Text = fmt.Sprintf("Error parsing challengeID: %v, err:%v", query.Data, err)
		return msg
	}
	msg.ReplyMarkup = tgbotapi.ForceReply{
		ForceReply: true,
	}

	msg.Text = fmt.Sprintf(
		"Create a goal for yourself in this Challenge\nInput target:")

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
