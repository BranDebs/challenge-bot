package condition

import (
	"encoding/json"
	"fmt"

	"github.com/BranDebs/challenge-bot/internal/domain/condition/operator"
	"github.com/BranDebs/challenge-bot/internal/domain/entry"
)

// EmptyCondition represents an uninitialised Condition.
var EmptyCondition = Condition{}

/*
Condition represents the condition in a Challenge.
It is represented in JSON in the following way:
  [
    {
      name: "BodyFat",
      kind: "float",
      entry: "18.4",
      operator: "eq"
    },
    {
      name: "GymWeekly",
      kind: "boolean",
      entry: "true",
      operator: "neq"
    }
  ]
*/
type Condition struct {
	entry.Entry
	Operator operator.Operator `json:"operator"`
}

// New validates and initialises a new Condition.
func New(name string, kind entry.Kind, val entry.Value, operator operator.Operator) (Condition, error) {
	if len(name) == 0 || !kind.Valid() || val.Valid(kind) {
		return EmptyCondition, ErrValidation{
			kind:  kind,
			value: val,
		}
	}

	return Condition{
		Entry: entry.Entry{
			Name:  name,
			Kind:  kind,
			Value: val,
		},
		Operator: operator,
	}, nil
}

func FromJSON(data []byte) ([]*Condition, error) {
	conditions := make([]*Condition, 0)

	if err := json.Unmarshal(data, &conditions); err != nil {
		return nil, fmt.Errorf("unable to unmarshal condition err: %w", err)
	}

	return conditions, nil
}
