package goal

import (
	"context"

	"github.com/BranDebs/challenge-bot/command/base"
	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"
)

type findGoalParams struct {
	challengeID uint64
	userID      uint64
}

type FindGoalCommand struct {
	formatter Formatter
	logic     Logic
	parser    Parser
	msg       base.MsgData
}

func (c FindGoalCommand) Execute(ctx context.Context) (string, error) {
	params, err := c.parser.ParseFindGoal(ctx, c.msg)
	if err != nil {
		return err.Error(), err
	}

	goalObj, err := c.logic.FindGoal(ctx, *params)
	if err != nil {
		return err.Error(), err
	}

	return c.formatter.FormatFind(ctx, goalObj, c.msg.UserID), nil
}

func NewFindGoalCommand(msg base.MsgData, handler logic.Handler, validator validator.Validator) base.Command {
	return &FindGoalCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
