package operator

import (
	"fmt"
	"strconv"
	"strings"
)

type Operator uint8

const (
	Unknown Operator = iota
	Equal
	NotEqual
	LessThan
	LessThanEqual
	GreaterThan
	GreaterThanEqual
)

func Valid(opr Operator) bool {
	return opr > Unknown && opr <= GreaterThanEqual
}

// String returns string representation of a Kind.
func (opr Operator) String() string {
	switch opr {
	case Equal:
		return "eq"
	case NotEqual:
		return "neq"
	case LessThan:
		return "lt"
	case LessThanEqual:
		return "lte"
	case GreaterThan:
		return "gt"
	case GreaterThanEqual:
		return "gte"
	default:
		return "unknown"
	}
}

func (o *Operator) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return fmt.Errorf("unable to unquote string err: %w", err)
	}

	*o = FromString(s)
	return nil
}

func FromString(s string) Operator {
	switch strings.ToLower(s) {
	case "eq":
		return Equal
	case "neq":
		return NotEqual
	case "lt":
		return LessThan
	case "lte":
		return LessThanEqual
	case "gt":
		return GreaterThan
	case "gte":
		return GreaterThanEqual
	default:
		return Unknown
	}
}
