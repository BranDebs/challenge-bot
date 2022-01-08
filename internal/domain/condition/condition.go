package condition

// EmptyCondition represents an uninitialised Condition.
var EmptyCondition = Condition{}

// Condition represents the condition in a Challenge.
type Condition struct {
	kind  Kind
	value Value
}

// New validates and initialises a new Condition.
func New(kind Kind, val Value) (Condition, error) {
	if !kind.Valid() || val.Valid(kind) {
		return EmptyCondition, ErrValidation{
			kind:  kind,
			value: val,
		}
	}

	return Condition{
		kind:  kind,
		value: val,
	}, nil
}

// Kind returns what kind of Condition this is.
func (c Condition) Kind() Kind {
	return c.kind
}

// Value returns the value of the Condition.
func (c Condition) Value() Value {
	return c.value
}
