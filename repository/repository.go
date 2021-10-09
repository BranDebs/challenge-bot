package repository

import (
	"context"

	"github.com/BranDebs/challenge-bot/challenge"
)

type Filters map[string]interface{}

type Repository interface {
	CreateChallenge(ctx context.Context, challenge *challenge.Challenge) error
	FindChallenge(ctx context.Context, id uint64) (*challenge.Challenge, error)
	ListChallenges(ctx context.Context, filters Filters, offset, limit uint64) ([]*challenge.Challenge, error)

	CreateUser(ctx context.Context, user *challenge.User) error
	FindUser(ctx context.Context, id uint64) (*challenge.User, error)
	ListUsers(ctx context.Context, filter Filters, offset, limit uint64) ([]*challenge.User, error)

	CreateGoal(ctx context.Context, goal *challenge.Goal) error
	FindGoal(ctx context.Context, id uint64) (*challenge.Goal, error)
	ListGoals(ctx context.Context, filter Filters, offset, limit uint64) ([]*challenge.Goal, error)

	CreateProgress(ctx context.Context, progress *challenge.Progress) error
	FindProgress(ctx context.Context, id uint64) (*challenge.Progress, error)
	ListProgress(ctx context.Context, filter Filters, offset, limit uint64) ([]*challenge.Progress, error)
}
