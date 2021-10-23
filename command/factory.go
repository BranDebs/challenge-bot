package command

import (
	"errors"
	"strings"

	"github.com/BranDebs/challenge-bot/command/progress"

	"github.com/BranDebs/challenge-bot/command/goal"

	"github.com/BranDebs/challenge-bot/command/help"

	"github.com/BranDebs/challenge-bot/command/challenge"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command/base"
)

type Invoker func(msg base.MsgData, handler logic.Handler, validator validator.Validator) base.Command

type Factory struct {
	invokers []Invoker
}

func (f Factory) GetCommand(msg base.MsgData, handler logic.Handler, validator validator.Validator) (base.Command, error) {
	msgTokens := strings.Fields(msg.Msg)
	if len(msgTokens) == 0 {
		return nil, errors.New("no command provided")
	}

	for _, invoker := range f.invokers {
		cmd := invoker(msg, handler, validator)
		if cmd != nil {
			return cmd, nil
		}
	}
	return nil, errors.New("invalid command")
}

func NewFactory() *Factory {
	return &Factory{
		invokers: []Invoker{
			challenge.ChallengeCommandInvoker,
			help.HelpCommandInvoker,
			goal.GoalCommandInvoker,
			progress.ProgressCommandInvoker,
		},
	}
}
