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
	StaticMainChallangePageKeyboard() tgbotapi.ReplyKeyboardMarkup
	GetChallengesKeyboard(challenges []model.Challenge) tgbotapi.InlineKeyboardMarkup
	GetChallengeActionKeyboard(challengeID uint64, userID uint64) (tgbotapi.InlineKeyboardMarkup, error)
}

type KeyboardProviderImpl struct {
}

func NewKeyboardProvider() KeyboardProvider {
	return KeyboardProviderImpl{}
}

func (k KeyboardProviderImpl) GetChallengeActionKeyboard(challengeID uint64, userID uint64) (tgbotapi.InlineKeyboardMarkup, error) {
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

// StaticMainChallangePageKeyboard returns the main page keyboard with static fields
func (k KeyboardProviderImpl) StaticMainChallangePageKeyboard() tgbotapi.ReplyKeyboardMarkup {
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
func (k KeyboardProviderImpl) GetChallengesKeyboard(challenges []model.Challenge) tgbotapi.InlineKeyboardMarkup {
	challenges = testChallenges
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

var testChallenges = []model.Challenge{
	{
		ID:          1,
		Name:        "challenge1",
		Description: "hi i am a fun challenge woooooooo",
		UserIDs:     nil,
		StartDate:   1633838675,
		EndDate:     1633839675,
	},
	{
		ID:          2,
		Name:        "challenge2",
		Description: "hi i am a fun challengee",
		UserIDs:     nil,
		StartDate:   1633828675,
		EndDate:     1633838675,
	},
	{
		ID:          3,
		Name:        "challenge3",
		Description: "hi i am a fun challengeee",
		UserIDs:     nil,
		StartDate:   1633838675,
		EndDate:     1633839675,
	},
}

func RemoveInlineKeyboard(chatID int64, messageID int) tgbotapi.EditMessageReplyMarkupConfig {
	return tgbotapi.NewEditMessageReplyMarkup(chatID, messageID, tgbotapi.InlineKeyboardMarkup{})
}
