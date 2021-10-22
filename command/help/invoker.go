package help

import (
	"strings"

	common "github.com/BranDebs/challenge-bot/command/common"

	"github.com/BranDebs/challenge-bot/command/model"
	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"
)

type helpCommand uint

const (
	help = helpCommand(iota)
)

func (c helpCommand) String() string {
	switch c {
	case help:
		return "/help"
	}
	return ""
}

func HelpCommandInvoker(msg model.MsgData, handler logic.Handler, validator validator.Validator) common.Command {
	msgTokens := strings.Fields(msg.Msg)
	switch msgTokens[0] {
	case help.String():
		return NewHelpCommand()
	default:
		return nil
	}
}
