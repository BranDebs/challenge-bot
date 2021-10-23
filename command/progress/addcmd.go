package progress

import (
	"context"

	"github.com/BranDebs/challenge-bot/command/base"
	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"
)

type addProgressParams struct {
	userID        uint64
	challengeID   uint64
	schema        []byte
	dateTimestamp uint64
}

type AddProgressCommand struct {
	formatter Formatter
	logic     Logic
	parser    Parser
	msg       base.MsgData
}

func (c AddProgressCommand) Execute(ctx context.Context) (string, error) {
	params, err := c.parser.ParseAddProgress(ctx, c.msg)
	if err != nil {
		return err.Error(), err
	}

	err = c.logic.AddProgress(ctx, *params)
	if err != nil {
		return err.Error(), err
	}

	return c.formatter.FormatAdd(ctx), nil
}

func NewAddProgressCommand(msg base.MsgData, handler logic.Handler, validator validator.Validator) base.Command {
	return &AddProgressCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
