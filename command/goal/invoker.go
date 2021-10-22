package goal

import (
	"strings"

	"github.com/BranDebs/challenge-bot/command/common"
	"github.com/BranDebs/challenge-bot/command/model"
	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"
)

type goalCommand uint

const (
	create = goalCommand(iota)
	detail
)

func (c goalCommand) String() string {
	switch c {
	case create:
		return "/createg"
	case detail:
		return "/gdetail"
	}
	return ""
}

func GoalCommandInvoker(msg model.MsgData, handler logic.Handler, validator validator.Validator) common.Command {
	msgTokens := strings.Fields(msg.Msg)
	switch msgTokens[0] {
	case create.String():
		return NewCreateGoalCommand(msg, handler, validator)
	case detail.String():
		return NewFindGoalCommand(msg, handler, validator)
	default:
		return nil
	}
}
