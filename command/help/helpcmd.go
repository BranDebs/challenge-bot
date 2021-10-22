package help

import (
	"context"

	"github.com/BranDebs/challenge-bot/command"
)

const (
	helpText = `
*1. createChallenge* \n
/createc name description enddate_in_YYYY_MM_DD schema\n
e.g. /createc 'my challenge name' 'lose fat' 2021-11-25 '{"weight": int64}'\n

*2. listChallenge*\n*
/listc\n

*3. challengeDetail*
/cdetail challengeID\n
e.g. /cdetail 123\n

*4. joinChallenge*
/joinc challengeID\n
e.g. /joinc 123\n

*5. createGoal*
/createg challengeID goalSchema\n
e.g. /createg 123 '{"weight": 50}'\n

*6. goalDetail*
/gdetail challengeID\n
e.g. /gdetail 123\n

*7. addProgress*
/addp challengeID progressSchema\n
e.g. /addp 123 '{"weight": 51}'\n

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

func NewHelpCommand() command.Command {
	return &HelpCommand{}
}
