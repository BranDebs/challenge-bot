package progress

import "github.com/BranDebs/challenge-bot/internal/domain/condition"

var EmptyProgress = &Progress{}

type Progress struct {
	ID          uint64
	UserID      uint64
	ChallengeID uint64
	Values      []condition.Condition
	IsGoal      bool
	UpdatedAt   uint64
}

func New(id uint64, userID uint64, challengeID uint64, data []byte, isGoal bool, updatedAt uint64) (*Progress, error) {
	if id == 0 || userID == 0 || challengeID == 0 {
		return EmptyProgress, ErrInvalidProgress
	}

	// make data into conditions
	values := make([]condition.Condition, 0)

	return &Progress{
		ID:          id,
		UserID:      userID,
		ChallengeID: challengeID,
		Values:      values,
		IsGoal:      isGoal,
		UpdatedAt:   updatedAt,
	}, nil
}
