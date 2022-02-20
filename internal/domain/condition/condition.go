package condition

import (
	"github.com/BranDebs/challenge-bot/internal/domain/condition/kind"
	"github.com/BranDebs/challenge-bot/internal/domain/condition/operator"
	"github.com/BranDebs/challenge-bot/internal/domain/condition/value"
)

// EmptyCondition represents an uninitialised Condition.
var EmptyCondition = Condition{}

// Condition represents the condition in a Challenge.
type Condition struct {
	name     string
	kind     kind.Kind
	value    value.Value
	operator operator.Operator
}

// New validates and initialises a new Condition.
func New(name string, kind kind.Kind, val value.Value, operator operator.Operator) (Condition, error) {
	if len(name) == 0 || !kind.Valid() || val.Valid(kind) {
		return EmptyCondition, ErrValidation{
			kind:  kind,
			value: val,
		}
	}

	return Condition{
		name:     name,
		kind:     kind,
		value:    val,
		operator: operator,
	}, nil
}

// Kind returns what kind of Condition this is.
func (c Condition) Kind() kind.Kind {
	return c.kind
}

// Value returns the value of the Condition.
func (c Condition) Value() value.Value {
	return c.value
}

// Operator returns the operator used for comparing the condition
func (c Condition) Operator() operator.Operator {
	return c.operator
}
