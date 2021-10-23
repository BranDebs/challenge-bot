package challenge

import (
	"context"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command/base"
)

type findChallengeParams struct {
	challengeID uint64
}

type FindChallengeCommand struct {
	formatter Formatter
	logic     Logic
	parser    Parser
	msg       base.MsgData
}

func (c FindChallengeCommand) Execute(ctx context.Context) (string, error) {
	params, err := c.parser.ParseFindChallenge(ctx, c.msg)
	if err != nil {
		return err.Error(), err
	}

	challengeObj, err := c.logic.FindChallenge(ctx, *params)
	if err != nil {
		return err.Error(), err
	}

	return c.formatter.FormatFind(ctx, challengeObj, c.msg.UserID), nil
}

func NewFindChallengeCommand(msg base.MsgData, handler logic.Handler, validator validator.Validator) base.Command {
	return &FindChallengeCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
