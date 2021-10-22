package challenge

import (
	"context"

	common "github.com/BranDebs/challenge-bot/command/common"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command/model"
)

type findChallengeParams struct {
	challengeID uint64
}

type FindChallengeCommand struct {
	formatter Formatter
	logic     Logic
	parser    Parser
	msg       model.MsgData
}

func (c FindChallengeCommand) Execute(ctx context.Context) (string, error) {
	params, err := c.parser.ParseFindChallenge(ctx, c.msg)
	if err != nil {
		return err.Error(), err
	}

	challengeObj, err := c.logic.FindChallenge(ctx, *params)
	if err != nil {
		return err.Error(), err
	}

	return c.formatter.FormatFind(ctx, challengeObj, c.msg.UserID), nil
}

func NewFindChallengeCommand(msg model.MsgData, handler logic.Handler, validator validator.Validator) common.Command {
	return &FindChallengeCommand{
		formatter: NewFormatter(),
		logic:     NewLogic(handler),
		parser:    NewParser(validator),
		msg:       msg,
	}
}
