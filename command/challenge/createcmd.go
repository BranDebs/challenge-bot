package challenge

import (
	"context"

	"github.com/BranDebs/challenge-bot/command/base"
	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"
)

type createChallangeParams struct {
	id          uint64
	name        string
	userID      uint64
	startDate   uint64
	endDate     uint64
	description string
	schema      []byte
}

type CreateChallengeCommand struct {
	formatter Formatter
	logic     Logic
	parser    Parser
	msg       base.MsgData
}

func (c CreateChallengeCommand) Execute(ctx context.Context) (string, error) {
	params, err := c.parser.ParseCreateChallenge(ctx, c.msg)
	if err != nil {
		return err.Error(), err
	}

	challengeObj, err := c.logic.CreateChallenge(ctx, *params)
	if err != nil {
		return err.Error(), err
	}

	return c.formatter.FormatCreate(ctx, challengeObj), nil
}

func NewCreateChallengeCommand(msg base.MsgData, handler logic.Handler, validator validator.Validator) base.Command {
	return &CreateChallengeCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
