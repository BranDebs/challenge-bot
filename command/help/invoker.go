package help

import (
	"strings"

	"github.com/BranDebs/challenge-bot/command"
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

func HelpCommandInvoker(msg model.Msg, handler logic.Handler, validator validator.Validator) command.Command {
	msgTokens := strings.Fields(msg.Msg)
	switch msgTokens[0] {
	case help.String():
		return NewHelpCommand()
	default:
		return nil
	}
}
