package challenge

import (
	"strings"

	"github.com/BranDebs/challenge-bot/logic"
	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command/base"
)

type challengeCommand uint

const (
	create = challengeCommand(iota)
	list
	detail
	join
)

func (c challengeCommand) String() string {
	switch c {
	case create:
		return "/createc"
	case list:
		return "/listc"
	case detail:
		return "/cdetail"
	case join:
		return "/joinc"
	}
	return ""
}

func ChallengeCommandInvoker(msg base.MsgData, handler logic.Handler, validator validator.Validator) base.Command {
	msgTokens := strings.Fields(msg.Msg)
	switch msgTokens[0] {
	case create.String():
		return NewCreateChallengeCommand(msg, handler, validator)
	case list.String():
		return NewListChallengeCommand(msg, handler, validator)
	case detail.String():
		return NewFindChallengeCommand(msg, handler, validator)
	case join.String():
		return NewJoinChallengeCommand(msg, handler, validator)
	default:
		return nil
	}
}
