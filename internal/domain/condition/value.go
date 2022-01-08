package condition

import (
	"strconv"
	"strings"
)

// Value represents the value of a condition
type Value string

// Valid returns true if the value is valid.
// A valid Value is based off the Kind types.
func (v Value) Valid(kind Kind) bool {
	if !kind.Valid() {
		return false
	}

	s := strings.ToLower(string(v))

	switch kind {
	case Boolean:
		_, err := strconv.ParseBool(s)
		return err == nil
	case Integer:
		_, err := strconv.ParseInt(s, 10, 64)
		return err == nil
	case Float:
		_, err := strconv.ParseFloat(s, 32)
		return err == nil
	default:
		return false
	}
}
