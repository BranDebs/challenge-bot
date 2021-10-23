package goal

import (
	"context"

	"github.com/BranDebs/challenge-bot/command/base"
	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"
)

type createGoalParams struct {
	userID      uint64
	challengeID uint64
	schema      []byte
}

type CreateGoalCommand struct {
	formatter Formatter
	logic     Logic
	parser    Parser
	msg       base.MsgData
}

func (c CreateGoalCommand) Execute(ctx context.Context) (string, error) {
	params, err := c.parser.ParseCreateGoal(ctx, c.msg)
	if err != nil {
		return err.Error(), err
	}

	goalObj, err := c.logic.CreateGoal(ctx, *params)
	if err != nil {
		return err.Error(), err
	}

	return c.formatter.FormatCreate(ctx, goalObj), nil
}

func NewCreateGoalCommand(msg base.MsgData, handler logic.Handler, validator validator.Validator) base.Command {
	return &CreateGoalCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
