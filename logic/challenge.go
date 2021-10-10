package logic

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/BranDebs/challenge-bot/model"
	"github.com/BranDebs/challenge-bot/repository"
)

var (
	ErrInvalidChallenge = errors.New("invalid challenge")
)

type ChallengeHandler interface {
	CreateChallenge(ctx context.Context, challenge *model.Challenge) error
}

type challengeHandler struct {
	repo repository.Challenge
}

func (ch *challengeHandler) CreateChallenge(ctx context.Context, challenge *model.Challenge) error {
	if challenge == nil {
		return fmt.Errorf("%w: %+v", ErrInvalidChallenge, challenge)
	}
	log.Printf("Creating challenge: %+v", challenge)

	if err := ch.repo.CreateChallenge(ctx, challenge); err != nil {
		return err
	}

	return nil
}
