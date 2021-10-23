package progress

import (
	"context"
	"errors"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/model"
)

type Logic interface {
	AddProgress(ctx context.Context, params addProgressParams) error
	ListProgress(ctx context.Context, params listProgressParams) ([]*model.Progress, error)
}

var (
	ErrFailedToAddProgress = errors.New("failed to create progress")
)

type progressLogic struct {
	pHandler logic.ProgressHandler
}

func (p progressLogic) AddProgress(ctx context.Context, params addProgressParams) error {
	_, isSuccess, err := p.pHandler.CreateProgress(ctx, &model.Progress{
		UserID:      params.userID,
		ChallengeID: params.challengeID,
		Value:       params.schema,
		Date:        params.dateTimestamp,
	})

	if err != nil {
		return err
	}
	if !isSuccess {
		return ErrFailedToAddProgress
	}
	return nil
}

func (p progressLogic) ListProgress(ctx context.Context, params listProgressParams) ([]*model.Progress, error) {
	return p.pHandler.ListProgress(ctx, params.challengeID, params.userID)
}

func NewLogic(pHandler logic.ProgressHandler) Logic {
	return &progressLogic{pHandler: pHandler}
}
