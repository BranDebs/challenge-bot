package condition

import (
	"errors"
	"fmt"

	"github.com/BranDebs/challenge-bot/internal/domain/condition/operator"
	"github.com/BranDebs/challenge-bot/internal/domain/entry"
)

type Comparer func(goal *Condition, progress *Condition) (bool, error)

func EqualComparer(goal *Condition, progress *Condition) (bool, error) {
	if goal == nil || progress == nil {
		return false, errors.New("invalid comparer input")
	}

	if goal.Operator != operator.Equal {
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

	if goal.Operator != operator.NotEqual {
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

	if goal.Operator != operator.LessThan {
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

	if goal.Operator != operator.LessThanEqual {
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

	if goal.Operator != operator.GreaterThan {
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

	if goal.Operator != operator.GreaterThan {
		return false, errors.New("invalid operator")
	}

	cmp, err := compare(goal, progress)
	if err != nil {
		return false, fmt.Errorf("failed to compare >= err: %w", err)
	}

	return cmp == 1 || cmp == 0, nil
}

func compare(goal, progress *Condition) (int, error) {
	switch goal.Kind {
	case entry.Boolean:
		gBool, gErr := entry.ParseBool(goal.Value)
		pBool, pErr := entry.ParseBool(progress.Value)
		if gErr != nil || pErr != nil {
			return 0, errors.New("failed to compare")
		}
		if gBool == pBool {
			return 0, nil
		}
		return -1, nil

	case entry.Integer:
		gInt, gErr := entry.ParseInt(goal.Value)
		pInt, pErr := entry.ParseInt(progress.Value)
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

	case entry.Float:
		gFloat, gErr := entry.ParseFloat(goal.Value)
		pFloat, pErr := entry.ParseFloat(progress.Value)
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
