package challenge

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/BranDebs/challenge-bot/command/util"

	"github.com/BranDebs/challenge-bot/validator"

	"github.com/BranDebs/challenge-bot/command/model"
)

type Parser interface {
	ParseCreateChallenge(ctx context.Context, msg model.MsgData) (*createChallangeParams, error)
	ParseListChallenges(ctx context.Context, msg model.MsgData) (*listChallangeParams, error)
	ParseFindChallenge(ctx context.Context, msg model.MsgData) (*findChallengeParams, error)
	ParseJoinChallenge(ctx context.Context, msg model.MsgData) (*joinChallengeParams, error)
}

const (
	createNumTokens = 5
	findNumTokens   = 2
	joinNumTokens   = 2
	layoutISO       = "2006-01-02"

	invalidDateStringErr = "invalid date string provided"
)

type parser struct {
	validator validator.Validator
}

// Create challenge format: /createc name description enddate-YYYY-MM-DD schema
// e.g. /createc 'my challenge name' 'lose fat' 2021-11-11 '{"weight": float}'
func (p parser) ParseCreateChallenge(ctx context.Context, msg model.MsgData) (*createChallangeParams, error) {
	tokens := util.GetTokens(msg.Msg)
	if !util.IsCorrectNumTokens(tokens, createNumTokens) {
		return nil, errors.New(util.InvalidTokenCountErr)
	}
	currTimestamp := uint64(time.Now().Unix())

	name := tokens[1]
	description := tokens[2]
	endDateString := tokens[3]
	schema := tokens[4]

	if err := p.validator.ValidateSchemaString(schema); err != nil {
		return nil, err
	}

	endDateTimestamp, err := parseDateString(endDateString)
	if err != nil {
		return nil, err
	}
	if err := p.validator.ValidateEndDateString(currTimestamp, endDateTimestamp); err != nil {
		return nil, err
	}

	return &createChallangeParams{
		name:        name,
		userID:      msg.UserID,
		startDate:   currTimestamp,
		endDate:     endDateTimestamp,
		description: description,
		schema:      []byte(schema),
	}, nil
}

func parseDateString(date string) (uint64, error) {
	t, err := time.Parse(layoutISO, date)
	if err != nil {
		return 0, errors.New(invalidDateStringErr)
	}
	return uint64(t.Unix()), nil
}

// List challenge format: /listc
// e.g. /listc
func (p parser) ParseListChallenges(ctx context.Context, msg model.MsgData) (*listChallangeParams, error) {
	return &listChallangeParams{userID: msg.UserID}, nil
}

// Find challenge format: /cdetail challengeID
// e.g. /cdetail 123
func (p parser) ParseFindChallenge(ctx context.Context, msg model.MsgData) (*findChallengeParams, error) {
	tokens := util.GetTokens(msg.Msg)
	if !util.IsCorrectNumTokens(tokens, findNumTokens) {
		return nil, errors.New(util.InvalidTokenCountErr)
	}

	challengeIDString := tokens[1]
	if err := p.validator.ValidateID(challengeIDString); err != nil {
		return nil, err
	}

	challengeID, _ := strconv.ParseUint(challengeIDString, 10, 64)

	return &findChallengeParams{
		challengeID: challengeID,
	}, nil
}

// Join challenge format: /joinc challengeID
// e.g. /joinc 123
func (p parser) ParseJoinChallenge(ctx context.Context, msg model.MsgData) (*joinChallengeParams, error) {
	tokens := util.GetTokens(msg.Msg)
	if !util.IsCorrectNumTokens(tokens, joinNumTokens) {
		return nil, errors.New(util.InvalidTokenCountErr)
	}

	challengeIDString := tokens[1]
	if err := p.validator.ValidateID(challengeIDString); err != nil {
		return nil, err
	}

	challengeID, _ := strconv.ParseUint(challengeIDString, 10, 64)

	return &joinChallengeParams{
		challengeID: challengeID,
		userID:      msg.UserID,
	}, nil
}

func NewParser(validator validator.Validator) Parser {
	return &parser{validator: validator}
}
