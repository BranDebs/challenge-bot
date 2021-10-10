package ui

import (
	"fmt"

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
}

func NewMessage() Message {
	return MessageImpl{}
}

func (m MessageImpl) GetMainScreenMsg(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	msg.Text = "Get started on Challenges =)"
	msg.ReplyMarkup = getMainChallangePageKeyboard()
	return msg
}

func (m MessageImpl) GetAvailableChallengesMsg(msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	// TODO: Retrieve available challenges here
	board := GetChallengesKeyboard(nil)
	msg.Text = GetChallengesText(nil, All)
	msg.ParseMode = parseMode
	msg.ReplyMarkup = board
	return msg
}

func (m MessageImpl) GetUserChallengesMsg(msg tgbotapi.MessageConfig, userID int) tgbotapi.MessageConfig {
	// TODO: Retrieve user's enrolled challenges here
	board := GetChallengesKeyboard(nil)
	msg.Text = GetChallengesText(nil, User)
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
	msg.ReplyMarkup = challengeDashboardPageKeyboard()
	return msg
}
