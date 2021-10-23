package challenge

import (
	"context"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/model"
)

type Logic interface {
	CreateChallenge(ctx context.Context, params createChallangeParams) (*model.Challenge, error)
	ListChallenges(ctx context.Context, params listChallangeParams) ([]*model.Challenge, error)
	FindChallenge(ctx context.Context, params findChallengeParams) (*model.Challenge, error)
	JoinChallenge(ctx context.Context, params joinChallengeParams) (*model.Challenge, error)
}

type challengeLogic struct {
	cHandler logic.ChallengeHandler
}

func (l challengeLogic) CreateChallenge(ctx context.Context, params createChallangeParams) (*model.Challenge, error) {
	challengeStruct := &model.Challenge{
		Name:        params.name,
		UserIDs:     []uint64{params.userID},
		StartDate:   params.startDate,
		EndDate:     params.endDate,
		Description: params.description,
		Schema:      params.schema,
	}
	challenge, err := l.cHandler.CreateChallenge(ctx, challengeStruct)
	if err != nil {
		return nil, err
	}

	// TODO: Return struct
	return challenge, nil
}

func (l challengeLogic) ListChallenges(ctx context.Context, params listChallangeParams) ([]*model.Challenge, error) {
	challenges, err := l.cHandler.ListChallenges(ctx)
	if err != nil {
		return nil, err
	}

	return challenges, nil
}

func (l challengeLogic) FindChallenge(ctx context.Context, params findChallengeParams) (*model.Challenge, error) {
	return l.cHandler.FindChallenge(ctx, params.challengeID)
}

func (l challengeLogic) JoinChallenge(ctx context.Context, params joinChallengeParams) (*model.Challenge, error) {
	err := l.cHandler.JoinChallenge(ctx, params.challengeID, params.userID)
	if err != nil {
		return nil, err
	}

	return l.cHandler.FindChallenge(ctx, params.challengeID)
}

func NewLogic(cHandler logic.ChallengeHandler) Logic {
	return &challengeLogic{cHandler: cHandler}
}
