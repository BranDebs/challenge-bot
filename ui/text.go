package ui

import (
	"fmt"
	"strings"

	"github.com/BranDebs/challenge-bot/model"
)

func GetChallengesText(challenges []model.Challenge, challengeType ChallengeType) string {
	challenges = testChallenges
	challengesText := getChallengeText(challengeType)
	for i, challenge := range challenges {
		x := fmt.Sprintf("*%v\\) %v*\n Description: %v \n StartDate: %v\n EndDate: %v\n\n",
			i+1,
			challenge.Name,
			challenge.Description,
			formatTimestampToDate(int64(challenge.StartDate)),
			formatTimestampToDate(int64(challenge.EndDate)),
		)
		challengesText = challengesText + x
	}

	challengesText = strings.ReplaceAll(challengesText, "_", "\\_")
	challengesText = strings.ReplaceAll(challengesText, "-", "\\-")
	challengesText = strings.ReplaceAll(challengesText, "[", "\\[")
	challengesText = strings.ReplaceAll(challengesText, "`", "\\`")
	challengesText = strings.ReplaceAll(challengesText, "(", "\\(")

	return challengesText
}

func getChallengeText(challengeType ChallengeType) string {
	switch challengeType {
	case All:
		return "*Available Challenges:*\n"
	case User:
		return "*Your Challenges:*\n"
	default:
		return ""
	}
}