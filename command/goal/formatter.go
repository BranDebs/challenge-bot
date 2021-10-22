package goal

import (
	"context"
	"fmt"

	"github.com/BranDebs/challenge-bot/command/util"

	"github.com/BranDebs/challenge-bot/model"
)

type Formatter interface {
	FormatCreate(ctx context.Context, goal *model.Goal) string
	FormatFind(ctx context.Context, goal *model.Goal, userID uint64) string
}

type formatter struct{}

func (f formatter) FormatCreate(ctx context.Context, goal *model.Goal) string {
	formattedGoal := f.formatGoal(goal)
	return util.CleanMarkdownMsg(formattedGoal)
}

func (f formatter) FormatFind(ctx context.Context, goal *model.Goal, userID uint64) string {
	formattedGoal := f.formatGoal(goal)
	return util.CleanMarkdownMsg(formattedGoal)
}

func (f formatter) formatGoal(goal *model.Goal) string {
	return fmt.Sprintf("*Goal:*\n ChallengeID: %v \n Aim: %v \n",
		goal.ChallengeID,
		goal.Value,
	)
}

func NewFormatter() Formatter {
	return &formatter{}
}
