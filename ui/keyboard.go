package ui

import (
	"encoding/json"
	"strconv"

	"github.com/BranDebs/challenge-bot/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	ChallengeActionList = []string{
		AddProgress,
		ViewAllUsersProgress,
		ViewYourProgressTimeline,
	}
)

type ActionDetail struct {
	action      string
	challengeID uint64
	userID      uint64
}

type KeyboardProvider interface {
	StaticMainChallengePageKeyboard() tgbotapi.ReplyKeyboardMarkup
	GetChallengesKeyboard(challenges []*model.Challenge) tgbotapi.InlineKeyboardMarkup
	GetChallengeActionKeyboard(challengeID uint64, userID uint64) (tgbotapi.InlineKeyboardMarkup, error)
}

type keyboardProvider struct {
}

func NewKeyboardProvider() KeyboardProvider {
	return keyboardProvider{}
}

func (k keyboardProvider) GetChallengeActionKeyboard(challengeID uint64, userID uint64) (tgbotapi.InlineKeyboardMarkup, error) {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	for _, action := range ChallengeActionList {
		actionDetail := &ActionDetail{
			action:      action,
			challengeID: challengeID,
			userID:      userID,
		}
		actionDetailStr, err := json.Marshal(actionDetail)
		if err != nil {
			return tgbotapi.InlineKeyboardMarkup{}, err
		}

		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(action, string(actionDetailStr)))
	}

	buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(BackKeyword, BackKeyword))
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, button := range buttons {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(button))
	}
	return keyboard, nil
}

// StaticMainChallengePageKeyboard returns the main page keyboard with static fields
func (k keyboardProvider) StaticMainChallengePageKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(JoinAChallenge),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(ViewYourChallenges),
		),
	)
	keyboard.OneTimeKeyboard = true
	return keyboard
}

// GetChallengesKeyboard takes in a list of challenges and returns the InlineKeyboardMarkup
// with the challenge name and a Callback query with the challenge ID (in string).
func (k keyboardProvider) GetChallengesKeyboard(challenges []*model.Challenge) tgbotapi.InlineKeyboardMarkup {
	buttons := make([]tgbotapi.InlineKeyboardButton, 0)
	for _, challenge := range challenges {
		buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(challenge.Name, strconv.FormatUint(challenge.ID, 10)))
	}

	buttons = append(buttons, tgbotapi.NewInlineKeyboardButtonData(BackKeyword, BackKeyword))
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, button := range buttons {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard,
			tgbotapi.NewInlineKeyboardRow(button))
	}
	return keyboard
}

func RemoveInlineKeyboard(chatID int64, messageID int) tgbotapi.EditMessageReplyMarkupConfig {
	return tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, tgbotapi.InlineKeyboardMarkup{})
}
