package value

import (
	"errors"
	"strconv"
	"strings"

	"github.com/BranDebs/challenge-bot/internal/domain/condition/kind"
)

// Value represents the value of a condition
type Value string

// Valid returns true if the value is valid.
// A valid Value is based off the Kind types.
func (v Value) Valid(k kind.Kind) bool {
	if !k.Valid() {
		return false
	}

	s := strings.ToLower(string(v))

	switch k {
	case kind.Boolean:
		_, err := strconv.ParseBool(s)
		return err == nil
	case kind.Integer:
		_, err := strconv.ParseInt(s, 10, 64)
		return err == nil
	case kind.Float:
		_, err := strconv.ParseFloat(s, 32)
		return err == nil
	default:
		return false
	}
}

func ParseBool(v Value) (bool, error) {
	if !v.Valid(kind.Boolean) {
		return false, errors.New("invalid value")
	}

	return strconv.ParseBool(string(v))
}

func ParseInt(v Value) (int64, error) {
	if !v.Valid(kind.Integer) {
		return 0, errors.New("invalid value")
	}

	return strconv.ParseInt(string(v), 10, 64)
}

func ParseFloat(v Value) (float64, error) {
	if !v.Valid(kind.Float) {
		return 0, errors.New("invalid value")
	}

	return strconv.ParseFloat(string(v), 64)
}
