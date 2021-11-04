package goal

import (
	"context"
	"fmt"

	"github.com/BranDebs/challenge-bot/command/base"

	"github.com/BranDebs/challenge-bot/model"
)

type Formatter interface {
	FormatCreate(ctx context.Context, goal *model.Goal) string
	FormatFind(ctx context.Context, goal *model.Goal, userID uint64) string
}

type formatter struct{}

func (f formatter) FormatCreate(ctx context.Context, goal *model.Goal) string {
	formattedGoal := f.formatGoal(goal)
	return formattedGoal
}

func (f formatter) FormatFind(ctx context.Context, goal *model.Goal, userID uint64) string {
	formattedGoal := f.formatGoal(goal)
	return formattedGoal
}

func (f formatter) formatGoal(goal *model.Goal) string {
	schemaMapValue := base.FormatSchemaValue(goal.Value)

	goalStr := fmt.Sprintf("*Goal:*\n ChallengeID: %v \n",
		goal.ChallengeID,
	)
	for k, v := range schemaMapValue {
		goalStr = goalStr + fmt.Sprintf("%v: %v\n", k, v)
	}

	return goalStr
}

func NewFormatter() Formatter {
	return &formatter{}
}
