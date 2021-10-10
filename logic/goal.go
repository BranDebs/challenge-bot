package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/BranDebs/challenge-bot/model"
	"github.com/BranDebs/challenge-bot/repository"
)

var (
	ErrInvalidGoal = errors.New("invalid goal")
	ErrGoalExists  = errors.New("goal exists")
)

type GoalHandler interface {
	CreateGoal(ctx context.Context, goal *model.Goal) error
	FindGoal(ctx context.Context, challengeID, userID uint64) (*model.Goal, error)
}

type goalHandler struct {
	repo repository.Goal
}

func (gh goalHandler) CreateGoal(ctx context.Context, goal *model.Goal) error {
	if goal == nil {
		return fmt.Errorf("%w: %+v", ErrInvalidGoal, goal)
	}

	filters := repository.Filters{
		"challenge_id": goal.ChallengeID,
		"user_id":      goal.UserID,
	}

	goals, err := gh.repo.ListGoals(ctx, filters, repository.DefaultOffset, repository.DefaultLimit)
	if err != nil {
		return err
	}

	if len(goals) > 0 {
		return ErrGoalExists
	}

	if err := gh.repo.CreateGoal(ctx, goal); err != nil {
		return err
	}

	return nil
}

func (gh goalHandler) FindGoal(ctx context.Context, challengeID, userID uint64) (*model.Goal, error) {
	filters := repository.Filters{
		"challenge_id": challengeID,
		"user_id":      userID,
	}

	goals, err := gh.repo.ListGoals(ctx, filters, repository.DefaultOffset, repository.DefaultLimit)
	if err != nil {
		return nil, err
	}

	if len(goals) == 0 {
		return nil, errors.New("no data")
	}

	if len(goals) > 1 {
		return nil, ErrInvalidGoal
	}

	return goals[0], nil
}
