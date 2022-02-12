package challenge

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/BranDebs/challenge-bot/internal/domain/condition"
)

type ConditionTypes map[string]condition.Kind

func NewConditionTypes(data []byte) ConditionTypes {
	if len(data) == 0 {
		return nil
	}

	ct := make(ConditionTypes, len(data))
	if err := json.Unmarshal(data, &ct); err != nil {
		fmt.Printf("err: %v", err)
		return nil
	}

	return ct
}

func (ct ConditionTypes) ToBytes() ([]byte, error) {
	if len(ct) == 0 {
		return nil, errors.New("unable to serialise empty conditions")
	}

	m := make(map[string]interface{}, len(ct))

	return json.Marshal(m)
}
