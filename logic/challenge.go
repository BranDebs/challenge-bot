package logic

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/BranDebs/challenge-bot/common/slices"
	"github.com/BranDebs/challenge-bot/model"
	"github.com/BranDebs/challenge-bot/repository"
)

var (
	ErrInvalidChallenge = errors.New("invalid challenge")
)

type ChallengeHandler interface {
	CreateChallenge(ctx context.Context, challenge *model.Challenge) error
	ListChallenges(ctx context.Context) ([]*model.Challenge, error)
	FindChallenge(ctx context.Context, id uint64) (*model.Challenge, error)
	JoinChallenge(ctx context.Context, challengeID, userID uint64) error
}

type challengeHandler struct {
	repo repository.Challenge
}

func (ch challengeHandler) CreateChallenge(ctx context.Context, challenge *model.Challenge) error {
	if challenge == nil {
		return fmt.Errorf("%w: %+v", ErrInvalidChallenge, challenge)
	}
	log.Printf("Creating challenge: %+v", challenge)

	if err := ch.repo.CreateChallenge(ctx, challenge); err != nil {
		return err
	}

	return nil
}

func (ch challengeHandler) ListChallenges(ctx context.Context) ([]*model.Challenge, error) {
	log.Println("Listing challenges")

	challenges, err := ch.repo.ListChallenges(ctx, nil, repository.DefaultOffset, repository.DefaultLimit)
	if err != nil {
		return nil, err
	}

	return challenges, nil
}

func (ch challengeHandler) FindChallenge(ctx context.Context, id uint64) (*model.Challenge, error) {
	log.Println("Find challenge")

	challenge, err := ch.repo.FindChallenge(ctx, id)
	if err != nil {
		return nil, err
	}

	return challenge, nil
}

func (ch challengeHandler) JoinChallenge(ctx context.Context, challengeID, userID uint64) error {
	challenge, err := ch.repo.FindChallenge(ctx, challengeID)
	if err != nil {
		return err
	}

	if slices.Uint64Exists(userID, challenge.UserIDs...) {
		return nil
	}

	challenge.UserIDs = append(challenge.UserIDs, userID)

	if err = ch.repo.UpdateChallenge(ctx, challenge); err != nil {
		return err
	}

	return nil
}
