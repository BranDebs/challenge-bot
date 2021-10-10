package logic

import (
	"context"

	"github.com/BranDebs/challenge-bot/model"
)

type UserHandler interface {
	RegisterUser(ctx context.Context, user *model.User) error
	ListChallenges(ctx context.Context, userIDs ...uint64) ([]*model.Challenge, error)
}
