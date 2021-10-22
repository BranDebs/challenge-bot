package challenge

import (
	"context"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command"
	"github.com/BranDebs/challenge-bot/command/model"
)

type findChallengeParams struct {
	challengeID uint64
}

type FindChallengeCommand struct {
	formatter Formatter
	logic     Logic
	parser    Parser
	msg       model.Msg
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

func NewFindChallengeCommand(msg model.Msg, handler logic.Handler, validator validator.Validator) command.Command {
	return &FindChallengeCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
