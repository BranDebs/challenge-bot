package condition

import (
	"fmt"

	"github.com/BranDebs/challenge-bot/internal/domain/condition/kind"
	"github.com/BranDebs/challenge-bot/internal/domain/condition/value"
)

// ErrValidation contains the error for validation.
type ErrValidation struct {
	kind  kind.Kind
	value value.Value
}

func (err ErrValidation) Error() string {
	return fmt.Sprintf("invalid condition with kind: %s and value: %s", err.kind, err.value)
}
