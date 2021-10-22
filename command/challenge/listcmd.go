package challenge

import (
	"context"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command"
	"github.com/BranDebs/challenge-bot/command/model"
)

type listChallangeParams struct {
	userID uint64
}

type ListChallengeCommand struct {
	formatter Formatter
	logic     Logic
	parser    Parser
	msg       model.Msg
}

func (c ListChallengeCommand) Execute(ctx context.Context) (string, error) {
	params, err := c.parser.ParseListChallenges(ctx, c.msg)
	if err != nil {
		return err.Error(), err
	}

	challengesObj, err := c.logic.ListChallenges(ctx, *params)
	if err != nil {
		return err.Error(), err
	}

	return c.formatter.FormatList(ctx, challengesObj, c.msg.UserID), nil
}

func NewListChallengeCommand(msg model.Msg, handler logic.Handler, validator validator.Validator) command.Command {
	return &ListChallengeCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
