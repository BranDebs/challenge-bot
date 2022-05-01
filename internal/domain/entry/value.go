package entry

import (
	"errors"
	"strconv"
	"strings"
)

// Value represents the value of a condition
type Value string

// Valid returns true if the value is valid.
// A valid Value is based off the Kind types.
func (v Value) Valid(k Kind) bool {
	if !k.Valid() {
		return false
	}

	s := strings.ToLower(string(v))

	switch k {
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

func ParseBool(v Value) (bool, error) {
	if !v.Valid(Boolean) {
		return false, errors.New("invalid value")
	}

	return strconv.ParseBool(string(v))
}

func ParseInt(v Value) (int64, error) {
	if !v.Valid(Integer) {
		return 0, errors.New("invalid value")
	}

	return strconv.ParseInt(string(v), 10, 64)
}

func ParseFloat(v Value) (float64, error) {
	if !v.Valid(Float) {
		return 0, errors.New("invalid value")
	}

	return strconv.ParseFloat(string(v), 64)
}
