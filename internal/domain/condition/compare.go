package condition

import (
	"errors"
	"fmt"

	"github.com/BranDebs/challenge-bot/internal/domain/condition/kind"
	"github.com/BranDebs/challenge-bot/internal/domain/condition/operator"
	"github.com/BranDebs/challenge-bot/internal/domain/condition/value"
)

type Comparer func(goal *Condition, progress *Condition) (bool, error)

func EqualComparer(goal *Condition, progress *Condition) (bool, error) {
	if goal == nil || progress == nil {
		return false, errors.New("invalid comparer input")
	}

	if goal.operator != operator.Equal {
		return false, errors.New("invalid operator")
	}

	cmp, err := compare(goal, progress)
	if err != nil {
		return false, fmt.Errorf("failed to compare == err: %w", err)
	}

	return cmp == 0, nil
}

func NotEqualComparer(goal, progress *Condition) (bool, error) {
	if goal == nil || progress == nil {
		return false, errors.New("invalid comparer input")
	}

	if goal.operator != operator.NotEqual {
		return false, errors.New("invalid operator")
	}

	cmp, err := compare(goal, progress)
	if err != nil {
		return false, fmt.Errorf("failed to compare != err: %w", err)
	}

	return cmp != 0, nil
}

func LessThanComparer(goal, progress *Condition) (bool, error) {
	if goal == nil || progress == nil {
		return false, errors.New("invalid comparer input")
	}

	if goal.operator != operator.LessThan {
		return false, errors.New("invalid operator")
	}

	cmp, err := compare(goal, progress)
	if err != nil {
		return false, fmt.Errorf("failed to compare < err: %w", err)
	}

	return cmp == -1, nil
}

func LessThanEqualComparer(goal, progress *Condition) (bool, error) {
	if goal == nil || progress == nil {
		return false, errors.New("invalid comparer input")
	}

	if goal.operator != operator.LessThanEqual {
		return false, errors.New("invalid operator")
	}

	cmp, err := compare(goal, progress)
	if err != nil {
		return false, fmt.Errorf("failed to compare <= err: %w", err)
	}

	return cmp == -1 || cmp == 0, nil
}

func GreaterThanComparer(goal, progress *Condition) (bool, error) {
	if goal == nil || progress == nil {
		return false, errors.New("invalid comparer input")
	}

	if goal.operator != operator.GreaterThan {
		return false, errors.New("invalid operator")
	}

	cmp, err := compare(goal, progress)
	if err != nil {
		return false, fmt.Errorf("failed to compare > err: %w", err)
	}

	return cmp == 1, nil
}

func GreaterThanEqualComparer(goal, progress *Condition) (bool, error) {
	if goal == nil || progress == nil {
		return false, errors.New("invalid comparer input")
	}

	if goal.operator != operator.GreaterThan {
		return false, errors.New("invalid operator")
	}

	cmp, err := compare(goal, progress)
	if err != nil {
		return false, fmt.Errorf("failed to compare >= err: %w", err)
	}

	return cmp == 1 || cmp == 0, nil
}

func compare(goal, progress *Condition) (int, error) {
	switch goal.kind {
	case kind.Boolean:
		gBool, gErr := value.ParseBool(goal.value)
		pBool, pErr := value.ParseBool(progress.value)
		if gErr != nil || pErr != nil {
			return 0, errors.New("failed to compare")
		}
		if gBool == pBool {
			return 0, nil
		}
		return -1, nil

	case kind.Integer:
		gInt, gErr := value.ParseInt(goal.value)
		pInt, pErr := value.ParseInt(progress.value)
		if gErr != nil || pErr != nil {
			return 0, errors.New("failed to compare")
		}

		if gInt > pInt {
			return 1, nil
		}

		if gInt < pInt {
			return -1, nil
		}

		return 0, nil

	case kind.Float:
		gFloat, gErr := value.ParseFloat(goal.value)
		pFloat, pErr := value.ParseFloat(progress.value)
		if gErr != nil || pErr != nil {
			return 0, errors.New("failed to compare")
		}

		if gFloat > pFloat {
			return 1, nil
		}

		if gFloat < pFloat {
			return -1, nil
		}

		return 0, nil

	default:
		return 0, errors.New("invalid condition type")
	}
}
