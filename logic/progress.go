package logic

import (
	"context"

	"github.com/BranDebs/challenge-bot/model"
	"github.com/BranDebs/challenge-bot/repository"
)

type ProgressHandler interface {
	CreateProgress(ctx context.Context, progress *model.Progress) (bool, error)
	ListProgress(ctx context.Context, filters repository.Filters, offset, limit uint) ([]*model.Progress, error)
}
