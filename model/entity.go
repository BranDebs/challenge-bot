package model

type Challenge struct {
	ID          uint64
	Name        string
	UserIDs     []uint64
	StartDate   uint64
	EndDate     uint64
	Description string
	Schema      []byte
}

type User struct {
	ID       uint64
	Username string
}

type Goal struct {
	ID          uint64
	UserID      uint64
	ChallengeID uint64
	Value       []byte
}

type Progress struct {
	ID          uint64
	UserID      uint64
	ChallengeID uint64
	Value       []byte
	Date        uint64
}
