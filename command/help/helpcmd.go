package help

import (
	"context"

	"github.com/BranDebs/challenge-bot/command/base"
)

const (
	helpText = `
*1. createChallenge*
/createc name description enddate_in_YYYY_MM_DD schema
e.g. /createc 'my challenge name' 'lose fat' 2021-11-25 '{"weight": int64}'

*2. listChallenge*
/listc

*3. challengeDetail*
/cdetail challengeID
e.g. /cdetail 123

*4. joinChallenge*
/joinc challengeID
e.g. /joinc 123

*5. createGoal*
/createg challengeID goalSchema
e.g. /createg 123 '{"weight": 50}'

*6. goalDetail*
/gdetail challengeID
e.g. /gdetail 123

*7. addProgress*
/addp challengeID progressSchema
e.g. /addp 123 '{"weight": 51}'

*8. listProgress*
/listp challengeID
e.g. /listp 123
`
)

type HelpCommand struct {
}

func (c HelpCommand) Execute(ctx context.Context) (string, error) {
	return helpText, nil
}

func NewHelpCommand() base.Command {
	return &HelpCommand{}
}
