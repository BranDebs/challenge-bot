package repository

import (
	"context"

	"github.com/BranDebs/challenge-bot/model"
)

type Filters map[string]interface{}

type Challenge interface {
	CreateChallenge(ctx context.Context, challenge *model.Challenge) error
	FindChallenge(ctx context.Context, id uint64) (*model.Challenge, error)
	ListChallenges(ctx context.Context, filters Filters, offset, limit uint64) ([]*model.Challenge, error)
}

type User interface {
	CreateUser(ctx context.Context, user *model.User) error
	FindUser(ctx context.Context, id uint64) (*model.User, error)
	ListUsers(ctx context.Context, filter Filters, offset, limit uint64) ([]*model.User, error)
}

type Goal interface {
	CreateGoal(ctx context.Context, goal *model.Goal) error
	FindGoal(ctx context.Context, id uint64) (*model.Goal, error)
	ListGoals(ctx context.Context, filter Filters, offset, limit uint64) ([]*model.Goal, error)
}

type Progress interface {
	CreateProgress(ctx context.Context, progress *model.Progress) error
	FindProgress(ctx context.Context, id uint64) (*model.Progress, error)
	ListProgress(ctx context.Context, filter Filters, offset, limit uint64) ([]*model.Progress, error)
}

type Repository interface {
	Challenge

	User

	Goal

	Progress
}
