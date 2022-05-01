package kind

import (
	"fmt"
	"strconv"
	"strings"
)

// Kind represents the type of Condition.
type Kind uint8

// Currently supported kinds of Condition.
const (
	Unknown Kind = iota
	Boolean
	Integer
	Float
)

// Valid returns true if the Kind is supported.
func (k Kind) Valid() bool {
	return k > Unknown && k <= Float
}

// String returns string representation of a Kind.
func (k Kind) String() string {
	switch k {
	case Boolean:
		return "boolean"
	case Integer:
		return "integer"
	case Float:
		return "float"
	default:
		return "unknown"
	}
}

func (k *Kind) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return fmt.Errorf("unable to unquote string err: %w", err)
	}

	*k = FromString(s)
	return nil
}

func FromString(s string) Kind {
	switch strings.ToLower(s) {
	case "boolean":
		return Boolean
	case "integer":
		return Integer
	case "float":
		return Float
	default:
		return Unknown
	}
}
