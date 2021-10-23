package goal

import (
	"context"

	"github.com/BranDebs/challenge-bot/logic"

	"github.com/BranDebs/challenge-bot/model"
)

type Logic interface {
	CreateGoal(ctx context.Context, params createGoalParams) (*model.Goal, error)
	FindGoal(ctx context.Context, params findGoalParams) (*model.Goal, error)
}

type goalLogic struct {
	gHandler logic.GoalHandler
}

func (g goalLogic) CreateGoal(ctx context.Context, params createGoalParams) (*model.Goal, error) {
	goalObj := &model.Goal{
		UserID:      params.userID,
		ChallengeID: params.challengeID,
		Value:       params.schema,
	}
	goal, err := g.gHandler.CreateGoal(ctx, goalObj)
	if err != nil {
		return nil, err
	}

	return goal, nil
}

func (g goalLogic) FindGoal(ctx context.Context, params findGoalParams) (*model.Goal, error) {
	return g.gHandler.FindGoal(ctx, params.challengeID, params.userID)
}

func NewLogic(gHandler logic.GoalHandler) Logic {
	return &goalLogic{gHandler: gHandler}
}
