package challenge

import (
	"context"
	"fmt"

	"github.com/BranDebs/challenge-bot/command/base"

	"github.com/BranDebs/challenge-bot/model"
)

type Formatter interface {
	FormatList(ctx context.Context, challenges []*model.Challenge, userID uint64) string
	FormatCreate(ctx context.Context, challenge *model.Challenge) string
	FormatFind(ctx context.Context, challenge *model.Challenge, userID uint64) string
	FormatJoin(ctx context.Context, challenge *model.Challenge, userID uint64) string
}

type formatter struct {
}

func (f formatter) FormatList(ctx context.Context, challenges []*model.Challenge, userID uint64) string {
	challengesStr := ""
	for _, challenge := range challenges {
		challengesStr = challengesStr + f.formatChallenge(challenge, userID)
	}

	return challengesStr
}

func (f formatter) formatChallenge(challenge *model.Challenge, userID uint64) string {
	return fmt.Sprintf("*%v*\n ID: %v \n Description: %v \n StartDate: %v\n EndDate: %v\n Schema: %v\n Are you a Participant?: %v \n\n",
		challenge.Name,
		challenge.ID,
		challenge.Description,
		base.FormatTimestampToDate(int64(challenge.StartDate)),
		base.FormatTimestampToDate(int64(challenge.EndDate)),
		string(challenge.Schema),
		f.formatIsParticipant(challenge.UserIDs, userID),
	)
}

func (f formatter) formatIsParticipant(challengeUserIDs []uint64, userID uint64) string {
	for _, challengeUserID := range challengeUserIDs {
		if challengeUserID == userID {
			return "Yes"
		}
	}
	return "No"
}

func (f formatter) FormatCreate(ctx context.Context, challenge *model.Challenge) string {
	return "Successfully created challenge"
}

func (f formatter) FormatFind(ctx context.Context, challenge *model.Challenge, userID uint64) string {
	challengeStr := f.formatChallenge(challenge, userID)
	return challengeStr
}

func (f formatter) FormatJoin(ctx context.Context, challenge *model.Challenge, userID uint64) string {
	challengeStr := "Successfully joined challenge: \n"
	challengeStr = challengeStr + f.formatChallenge(challenge, userID)
	return challengeStr
}

func NewFormatter() Formatter {
	return &formatter{}
}
