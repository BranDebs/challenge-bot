package progress

import (
	"strings"

	"github.com/BranDebs/challenge-bot/command/base"
	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"
)

type progressCommand uint

const (
	add = progressCommand(iota)
	list
)

func (c progressCommand) String() string {
	switch c {
	case add:
		return "/addp"
	case list:
		return "/listp"
	}
	return ""
}

func ProgressCommandInvoker(msg base.MsgData, handler logic.Handler, validator validator.Validator) base.Command {
	msgTokens := strings.Fields(msg.Msg)
	switch msgTokens[0] {
	case add.String():
		return NewAddProgressCommand(msg, handler, validator)
	case list.String():
		return NewListProgressCommand(msg, handler, validator)
	default:
		return nil
	}
}
