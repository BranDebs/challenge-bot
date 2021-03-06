package goal

import (
	"context"
	"strconv"

	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command/base"
)

type Parser interface {
	ParseCreateGoal(ctx context.Context, msg base.MsgData) (*createGoalParams, error)
	ParseFindGoal(ctx context.Context, msg base.MsgData) (*findGoalParams, error)
}

const (
	createNumTokens = 3
	findNumTokens   = 2
)

type parser struct {
	validator validator.Validator
}

// Create goal format: /createg challengeID goalSchema
// e.g. /createg 123 {“weight”: 50}
func (p parser) ParseCreateGoal(ctx context.Context, msg base.MsgData) (*createGoalParams, error) {
	tokens := base.GetTokens(msg.Msg)
	if !base.IsCorrectNumTokens(tokens, createNumTokens) {
		return nil, base.ErrInvalidTokenCount
	}

	challengeIDString := tokens[1]
	goalSchemaString := tokens[2]
	if err := p.validator.ValidateID(challengeIDString); err != nil {
		return nil, err
	}
	if err := p.validator.ValidateSchemaString(goalSchemaString); err != nil {
		return nil, err
	}

	challengeID, _ := strconv.ParseUint(challengeIDString, 10, 64)

	return &createGoalParams{
		userID:      msg.UserID,
		challengeID: challengeID,
		schema:      []byte(goalSchemaString),
	}, nil
}

// Find goal format: /gdetail challengeID
// e.g. /gdetail 123
func (p parser) ParseFindGoal(ctx context.Context, msg base.MsgData) (*findGoalParams, error) {
	tokens := base.GetTokens(msg.Msg)
	if !base.IsCorrectNumTokens(tokens, findNumTokens) {
		return nil, base.ErrInvalidTokenCount
	}

	challengeIDString := tokens[1]
	challengeID, _ := strconv.ParseUint(challengeIDString, 10, 64)

	return &findGoalParams{
		userID:      msg.UserID,
		challengeID: challengeID,
	}, nil
}

func NewParser(validator validator.Validator) Parser {
	return &parser{validator: validator}
}
