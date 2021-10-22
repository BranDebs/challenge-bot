package challenge

import (
	"context"

	common "github.com/BranDebs/challenge-bot/command/common"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command/model"
)

type joinChallengeParams struct {
	challengeID uint64
	userID      uint64
}

type JoinChallengeCommand struct {
	formatter Formatter
	logic     Logic
	parser    Parser
	msg       model.MsgData
}

func (c JoinChallengeCommand) Execute(ctx context.Context) (string, error) {
	params, err := c.parser.ParseJoinChallenge(ctx, c.msg)
	if err != nil {
		return err.Error(), err
	}

	challengeObj, err := c.logic.JoinChallenge(ctx, *params)
	if err != nil {
		return err.Error(), err
	}

	return c.formatter.FormatJoin(ctx, challengeObj, c.msg.UserID), nil
}

func NewJoinChallengeCommand(msg model.MsgData, handler logic.Handler, validator validator.Validator) common.Command {
	return &JoinChallengeCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
