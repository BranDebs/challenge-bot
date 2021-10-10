package postgres

import (
	"github.com/BranDebs/challenge-bot/model"
)

type challengeEntity struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string
	Description string
	UserIDs     string
	StartDate   uint64
	EndDate     uint64
	Schema      []byte
}

func (e *challengeEntity) fromModel(challenge *model.Challenge) {
	e.ID = challenge.ID
	e.Name = challenge.Name
	e.Description = challenge.Description
	e.UserIDs = serializeUint64(challenge.UserIDs...)
	e.StartDate = challenge.StartDate
	e.EndDate = challenge.EndDate
	e.Schema = challenge.Schema
}

func (e challengeEntity) toModel() *model.Challenge {
	return &model.Challenge{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		UserIDs:     deserializeToUint64(e.UserIDs),
		StartDate:   e.StartDate,
		EndDate:     e.EndDate,
		Schema:      e.Schema,
	}
}

type goalEntity struct {
	ID          uint64 `gorm:"primaryKey"`
	UserID      uint64
	ChallengeID uint64
	Value       []byte
}

func (e *goalEntity) fromModel(goal *model.Goal) {
	e.ID = goal.ID
	e.ChallengeID = goal.ChallengeID
	e.UserID = goal.UserID
	e.Value = goal.Value
}

func (e goalEntity) toModel() *model.Goal {
	return &model.Goal{
		ID:          e.ID,
		ChallengeID: e.ChallengeID,
		UserID:      e.UserID,
		Value:       e.Value,
	}
}

type progressEntity struct {
	ID          uint64
	UserID      uint64
	ChallengeID uint64
	Value       []byte
	Date        uint64
}

func (e *progressEntity) fromModel(progress *model.Progress) {
	e.ID = progress.ID
	e.UserID = progress.UserID
	e.ChallengeID = progress.ChallengeID
	e.Value = progress.Value
	e.Date = progress.Date
}

func (e progressEntity) toModel() *model.Progress {
	return &model.Progress{
		ID:          e.ID,
		UserID:      e.UserID,
		ChallengeID: e.ChallengeID,
		Value:       e.Value,
		Date:        e.Date,
	}
}
