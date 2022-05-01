package condition

import (
	"github.com/BranDebs/challenge-bot/internal/domain/entry"
)

var units = map[string]entry.Kind{
	"kg": entry.Float,
	"g":  entry.Float,
}

func Units() map[string]entry.Kind {
	return units
}
