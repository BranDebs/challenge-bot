package challenge

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/BranDebs/challenge-bot/model"
)

type Formatter interface {
	FormatList(ctx context.Context, challenges []*model.Challenge, userID uint64) string
	FormatCreate(ctx context.Context, challenge *model.Challenge) string
	FormatFind(ctx context.Context, challenge *model.Challenge, userID uint64) string
	FormatJoin(ctx context.Context, challenge *model.Challenge, userID uint64) string
}

const (
	dateLayout = "02-01-2006 15:04:00"
)

type formatter struct {
}

func (f formatter) FormatList(ctx context.Context, challenges []*model.Challenge, userID uint64) string {
	challengesStr := ""
	for _, challenge := range challenges {
		challengesStr = challengesStr + f.formatChallenge(challenge, userID)
	}

	return f.cleanMarkdownMsg(challengesStr)
}

func (f formatter) formatChallenge(challenge *model.Challenge, userID uint64) string {
	return fmt.Sprintf("*%v*\n ID: %v \n Description: %v \n StartDate: %v\n EndDate: %v\n Are you a Participant?: %v \n\n",
		challenge.Name,
		challenge.ID,
		challenge.Description,
		formatTimestampToDate(int64(challenge.StartDate)),
		formatTimestampToDate(int64(challenge.EndDate)),
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

func (f formatter) cleanMarkdownMsg(msg string) string {
	msg = strings.ReplaceAll(msg, "_", "\\_")
	msg = strings.ReplaceAll(msg, "-", "\\-")
	msg = strings.ReplaceAll(msg, "[", "\\[")
	msg = strings.ReplaceAll(msg, "`", "\\`")
	msg = strings.ReplaceAll(msg, "(", "\\(")

	return msg
}

func formatTimestampToDate(timestamp int64) string {
	convertedTime := time.Unix(timestamp, 0)
	return convertedTime.Format(dateLayout)
}

func (f formatter) FormatCreate(ctx context.Context, challenge *model.Challenge) string {
	return "Successfully created challenge"
}

func (f formatter) FormatFind(ctx context.Context, challenge *model.Challenge, userID uint64) string {
	challengeStr := f.formatChallenge(challenge, userID)
	return f.cleanMarkdownMsg(challengeStr)
}

func (f formatter) FormatJoin(ctx context.Context, challenge *model.Challenge, userID uint64) string {
	challengeStr := "Successfully joined challenge: \n"
	challengeStr = challengeStr + f.formatChallenge(challenge, userID)
	return f.cleanMarkdownMsg(challengeStr)
}

func NewFormatter() Formatter {
	return &formatter{}
}
