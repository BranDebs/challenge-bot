package condition

import (
	"github.com/BranDebs/challenge-bot/internal/domain/condition/kind"
)

var units = map[string]kind.Kind{
	"kg": kind.Float,
	"g":  kind.Float,
}

func Units() map[string]kind.Kind {
	return units
}
