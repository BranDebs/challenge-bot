package condition

import (
	"fmt"

	"github.com/BranDebs/challenge-bot/internal/domain/entry"
)

// ErrValidation contains the error for validation.
type ErrValidation struct {
	kind  entry.Kind
	value entry.Value
}

func (err ErrValidation) Error() string {
	return fmt.Sprintf("invalid condition with kind: %s and entry: %s", err.kind, err.value)
}
