package progress

import (
	"context"

	"github.com/BranDebs/challenge-bot/command/base"
	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"
)

type listProgressParams struct {
	challengeID uint64
	userID      uint64
}

type ListProgressCommand struct {
	formatter Formatter
	logic     Logic
	parser    Parser
	msg       base.MsgData
}

func (c ListProgressCommand) Execute(ctx context.Context) (string, error) {
	params, err := c.parser.ParseListProgress(ctx, c.msg)
	if err != nil {
		return err.Error(), err
	}

	progressListObj, err := c.logic.ListProgress(ctx, *params)
	if err != nil {
		return err.Error(), err
	}

	return c.formatter.FormatList(ctx, progressListObj), nil
}

func NewListProgressCommand(msg base.MsgData, handler logic.Handler, validator validator.Validator) base.Command {
	return &ListProgressCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
