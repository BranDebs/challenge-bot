package challenge

import (
	"context"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command"
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
	msg       model.Msg
}

func (c JoinChallengeCommand) Execute(ctx context.Context) (string, error) {
	params, err := c.parser.ParseJoinChallenge(ctx, c.msg)
	if err != nil {
		return err.Error(), err
	}

	err = c.logic.JoinChallenge(ctx, *params)
	if err != nil {
		return err.Error(), err
	}

	return c.formatter.FormatJoin(ctx, nil, c.msg.UserID), nil
}

func NewJoinChallengeCommand(msg model.Msg, handler logic.Handler, validator validator.Validator) command.Command {
	return &JoinChallengeCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
