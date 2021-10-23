package challenge

import (
	"context"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command/base"
)

type joinChallengeParams struct {
	challengeID uint64
	userID      uint64
}

type JoinChallengeCommand struct {
	formatter Formatter
	logic     Logic
	parser    Parser
	msg       base.MsgData
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

func NewJoinChallengeCommand(msg base.MsgData, handler logic.Handler, validator validator.Validator) base.Command {
	return &JoinChallengeCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
