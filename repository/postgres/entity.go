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
