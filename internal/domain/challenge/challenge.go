package challenge

import (
	"errors"

	"github.com/BranDebs/challenge-bot/internal/domain/condition"
)

var (
	// ErrConditionTypesEmpty is raised when there is no condition type during creation or update of a challenge.
	ErrConditionTypesEmpty = errors.New("challenge must have at least a condition")
)

// EmptyChallenge represents an uninitialised challenge.
var EmptyChallenge = &Challenge{}

// Challenge represents the metadata of a challenge.
type Challenge struct {
	id                 uint64
	name               string
	description        string
	userIDs            []uint64
	conditions         []condition.Condition
	startDate, endDate uint64
}

// New validates and returns an initialised challenge.
func New(id uint64, name, description string, userIDs []uint64, data []byte, startDate, endDate uint64) (*Challenge, error) {
	if id == 0 || len(name) == 0 || len(userIDs) == 0 || endDate < startDate {
		return EmptyChallenge, ErrValidation{
			id:   id,
			name: name,
		}
	}

	cts := NewConditionTypes(data)
	if cts == nil {
		return EmptyChallenge, ErrConditionTypesEmpty
	}

	return &Challenge{
		name:           name,
		description:    description,
		userIDs:        userIDs,
		conditionTypes: cts,
		startDate:      startDate,
		endDate:        endDate,
	}, nil
}

func (c Challenge) ID() uint64 {
	return c.id
}

// Name returns the name of the Challenge.
func (c Challenge) Name() string {
	return c.name
}

// Description returns the description of the Challenge.
func (c Challenge) Description() string {
	return c.description
}

// UserIDs returns the userIDs that are in the Challenge.
func (c Challenge) UserIDs() []uint64 {
	return c.userIDs
}

func (c Challenge) ConditionTypes() ConditionTypes {
	return c.conditionTypes
}

// StartDate indicates the start of the Challenge.
func (c Challenge) StartDate() uint64 {
	return c.startDate
}

// EndDate indicates the end of the Challenge.
func (c Challenge) EndDate() uint64 {
	return c.endDate
}

// IsActive returns true if the Challenge is still active.
func (c Challenge) IsActive(now uint64) bool {
	return c.startDate <= now && now <= c.endDate
}

// HasUserID returns true if a userID is present in the Challenge.
func (c Challenge) HasUserID(userID uint64) bool {
	for _, id := range c.userIDs {
		if id == userID {
			return true
		}
	}
	return false
}

// AddUsers adds userIDs into the Challenge.
func (c *Challenge) AddUsers(userIDs ...uint64) error {
	for _, id := range userIDs {
		if id == 0 {
			return ErrUserID(id)
		}
		c.userIDs = append(c.userIDs, id)
	}
	return nil
}

// UpdateConditionTypes updates the condition types within the challenge.
// The update is a full update and not a delta.
func (c *Challenge) UpdateConditionTypes(data []byte) error {
	cts := NewConditionTypes(data)
	if cts == nil {
		return ErrConditionTypesEmpty
	}

	c.conditionTypes = cts

	return nil
}
