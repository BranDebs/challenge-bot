package challenge

import "fmt"

// ErrValidation contains the error for validation.
type ErrValidation struct {
	id   uint64
	name string
}

func (err ErrValidation) Error() string {
	return fmt.Sprintf("invalid challenge with id: %d and name: %s", err.id, err.name)
}

// ErrUserID contains the error for userID being invalid.
type ErrUserID uint64

func (err ErrUserID) Error() string {
	return fmt.Sprintf("invalid userID: %d", err)
}
