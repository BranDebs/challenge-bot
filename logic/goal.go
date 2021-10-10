package logic

import (
	"context"

	"github.com/BranDebs/challenge-bot/model"
)

type GoalHandler interface {
	CreateGoal(ctx context.Context, goal *model.Goal) error
	FindGoal(ctx context.Context, challengeID uint64) (*model.Goal, error)
}
