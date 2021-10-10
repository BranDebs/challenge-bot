package repository

import (
	"context"

	"github.com/BranDebs/challenge-bot/model"
)

type Filters map[string]interface{}

type Repository interface {
	CreateChallenge(ctx context.Context, challenge *model.Challenge) error
	FindChallenge(ctx context.Context, id uint64) (*model.Challenge, error)
	ListChallenges(ctx context.Context, filters Filters, offset, limit uint64) ([]*model.Challenge, error)

	CreateUser(ctx context.Context, user *model.User) error
	FindUser(ctx context.Context, id uint64) (*model.User, error)
	ListUsers(ctx context.Context, filter Filters, offset, limit uint64) ([]*model.User, error)

	CreateGoal(ctx context.Context, goal *model.Goal) error
	FindGoal(ctx context.Context, id uint64) (*model.Goal, error)
	ListGoals(ctx context.Context, filter Filters, offset, limit uint64) ([]*model.Goal, error)

	CreateProgress(ctx context.Context, progress *model.Progress) error
	FindProgress(ctx context.Context, id uint64) (*model.Progress, error)
	ListProgress(ctx context.Context, filter Filters, offset, limit uint64) ([]*model.Progress, error)
}
