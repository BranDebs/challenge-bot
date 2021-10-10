package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/BranDebs/challenge-bot/common/numbers"
	"github.com/BranDebs/challenge-bot/model"
	"github.com/BranDebs/challenge-bot/repository"
)

var (
	ErrInvalidProgress = errors.New("invalid progress")
	ErrCompareValues   = errors.New("unable to compare values")
)

type ProgressHandler interface {
	CreateProgress(ctx context.Context, progress *model.Progress) (bool, error)
	ListProgress(ctx context.Context, challengeID, userID uint64) ([]*model.Progress, error)
}

type progressHandler struct {
	cRepo repository.Challenge
	pRepo repository.Progress
	gRepo repository.Goal
}

func (ph progressHandler) CreateProgress(ctx context.Context, progress *model.Progress) (bool, error) {
	if progress == nil {
		return false, fmt.Errorf("%w: %v", ErrInvalidProgress, progress)
	}

	if err := ph.pRepo.CreateProgress(ctx, progress); err != nil {
		return false, nil
	}

	// Check if we met goal criteria
	filters := repository.Filters{
		"challenge_id": progress.ChallengeID,
		"user_id":      progress.UserID,
	}

	goals, err := ph.gRepo.ListGoals(ctx, filters, repository.DefaultOffset, repository.DefaultLimit)
	if err != nil {
		return false, err
	}

	if len(goals) > 1 {
		return false, ErrInvalidGoal
	}

	challenge, err := ph.cRepo.FindChallenge(ctx, goals[0].ChallengeID)
	if err != nil {
		return false, err
	}

	completed, err := completeChallenge(challenge.Schema, goals[0].Value, progress.Value)
	if err != nil {
		return false, err
	}

	return completed, nil
}

func (ph progressHandler) ListProgress(ctx context.Context, challengeID, userID uint64) ([]*model.Progress, error) {
	filters := repository.Filters{
		"challenge_id": challengeID,
		"user_id":      userID,
	}

	progress, err := ph.pRepo.ListProgress(ctx, filters, repository.DefaultOffset, repository.DefaultLimit)
	if err != nil {
		return nil, err
	}

	return progress, nil
}

func completeChallenge(schema, goalVal, progressVal []byte) (bool, error) {
	schemaJSON := make(map[string]interface{})
	goalJSON := make(map[string]interface{})
	progressJSON := make(map[string]interface{})

	if err := json.Unmarshal(schema, &schemaJSON); err != nil {
		return false, err
	}

	if err := json.Unmarshal(goalVal, &goalJSON); err != nil {
		return false, err
	}

	if err := json.Unmarshal(progressVal, &progressJSON); err != nil {
		return false, err
	}

	for k := range schemaJSON {
		cmp, err := compareValueType(k, schemaJSON, goalJSON, progressJSON)
		if err != nil {
			return false, ErrCompareValues
		}
		log.Printf("cmp: %v\n", cmp)
		if cmp < 0 {
			return false, nil
		}
	}

	return true, nil
}

func compareValueType(key string, schema, goal, progress map[string]interface{}) (int, error) {
	sv, ok := schema[key].(string)
	if !ok {
		return 0, ErrCompareValues
	}

	log.Printf("Comparing key: %+v schema: %+v goal: %+v progress: %+v\n", key, schema, goal, progress)

	switch sv {
	case "int64":
		log.Println("sv", sv)
		gv, ok := goal[key].(float64)
		if !ok {
			return 0, ErrCompareValues
		}

		log.Println("gv", gv)

		pv, ok := progress[key].(float64)
		if !ok {
			return 0, ErrCompareValues
		}

		return numbers.Uint64Compare(uint64(gv), uint64(pv)), nil
	}

	return 0, ErrCompareValues
}
