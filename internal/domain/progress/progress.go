package progress

type Progress struct {
	ID          uint64
	UserID      uint64
	ChallengeID uint64
	Value       []byte
	IsGoal      bool
	UpdatedAt   uint64
}
