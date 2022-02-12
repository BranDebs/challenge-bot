package condition

import "fmt"

// ErrValidation contains the error for validation.
type ErrValidation struct {
	kind  Kind
	value Value
}

func (err ErrValidation) Error() string {
	return fmt.Sprintf("invalid condition with kind: %s and value: %s", err.kind, err.value)
}
