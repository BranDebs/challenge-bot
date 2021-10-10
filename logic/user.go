package logic

import (
	"context"

	"github.com/BranDebs/challenge-bot/model"
	"github.com/BranDebs/challenge-bot/repository"
)

type UserHandler interface {
	RegisterUser(ctx context.Context, user *model.User) error
	ListChallenges(ctx context.Context, filters repository.Filters, offset, limit uint) ([]*model.Challenge, error)
}
