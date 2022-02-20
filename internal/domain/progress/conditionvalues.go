package progress

import (
	"encoding/json"

	"github.com/BranDebs/challenge-bot/internal/domain/condition"
)

type ConditionValues map[string]condition.Value

func NewConditionValues(data []byte) ConditionValues {
	if len(data) == 0 {
		return nil
	}
	ct := make(ConditionValues)
	if err := json.Unmarshal(data, &ct); err != nil {
		return nil
	}

	return ct
}
