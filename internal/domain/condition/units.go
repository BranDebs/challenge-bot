package condition

import (
	"github.com/BranDebs/challenge-bot/internal/domain/value"
)

var units = map[string]value.Kind{
	"kg": value.Float,
	"g":  value.Float,
}

func Units() map[string]value.Kind {
	return units
}
