package progress

import (
	"context"
	"fmt"

	"github.com/BranDebs/challenge-bot/command/base"

	"github.com/BranDebs/challenge-bot/model"
)

type Formatter interface {
	FormatAdd(ctx context.Context) string
	FormatList(ctx context.Context, progressList []*model.Progress) string
}

type formatter struct{}

func (f formatter) FormatAdd(ctx context.Context) string {
	return "Progress successfully added"
}

func (f formatter) FormatList(ctx context.Context, progressList []*model.Progress) string {
	var progressListStr string
	for _, progressObj := range progressList {
		progressStr := f.formatProgress(progressObj)
		progressListStr = progressStr + "\n"
	}

	return progressListStr
}

func (f formatter) formatProgress(progressObj *model.Progress) string {
	progressStr := fmt.Sprintf("*ChallengeID: %v*\n Date: %v\n ",
		progressObj.ChallengeID,
		base.FormatTimestampToDate(int64(progressObj.Date)),
	)

	schemaMapValue := base.FormatSchemaValue(progressObj.Value)
	for k, v := range schemaMapValue {
		progressStr = progressStr + fmt.Sprintf("%v: %v\n", k, v)
	}

	return progressStr
}

func NewFormatter() Formatter {
	return &formatter{}
}
