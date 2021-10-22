package command

import (
	"context"
	"errors"
	"strings"

	"github.com/BranDebs/challenge-bot/command/help"

	"github.com/BranDebs/challenge-bot/command/challenge"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command/model"
)

type Command interface {
	Execute(ctx context.Context) (string, error)
}
type Invoker func(msg model.Msg, handler logic.Handler, validator validator.Validator) Command

type Factory struct {
	invokers []Invoker
}

func (f Factory) GetCommand(msg model.Msg, handler logic.Handler, validator validator.Validator) (Command, error) {
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
	return &Factory{invokers: []Invoker{challenge.ChallengeCommandInvoker, help.HelpCommandInvoker}}
}
