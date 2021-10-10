package ui

import "time"

const (
	StartBot           = "/start"
	JoinAChallenge     = "Join a Challenge"
	ViewYourChallenges = "View your Challenges"

	AddProgress              = "Add Progress"
	ViewAllUsersProgress     = "View All users Progress"
	ViewYourProgressTimeline = "View you progress timeline"

	DateLayout = "02-01-2006 15:04:00"

	BackKeyword = "back"
	parseMode   = "MarkdownV2"
)

type ChallengeType int64

const (
	All  ChallengeType = 0
	User               = 1
)

func formatTimestampToDate(timestamp int64) string {
	convertedTime := time.Unix(timestamp, 0)
	return convertedTime.Format(DateLayout)
}
