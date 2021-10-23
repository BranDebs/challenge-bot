package progress

import (
	"context"
	"strconv"
	"time"

	"github.com/BranDebs/challenge-bot/command/base"
	"github.com/BranDebs/challenge-bot/validator"
)

type Parser interface {
	ParseAddProgress(ctx context.Context, msg base.MsgData) (*addProgressParams, error)
	ParseListProgress(ctx context.Context, msg base.MsgData) (*listProgressParams, error)
}

const (
	addNumTokens  = 3
	listNumTokens = 2
)

type parser struct {
	validator validator.Validator
}

func (p parser) ParseAddProgress(ctx context.Context, msg base.MsgData) (*addProgressParams, error) {
	tokens := base.GetTokens(msg.Msg)
	if !base.IsCorrectNumTokens(tokens, addNumTokens) {
		return nil, base.ErrInvalidTokenCount
	}

	challengeIDString := tokens[1]
	schemaString := tokens[2]
	if err := p.validator.ValidateID(challengeIDString); err != nil {
		return nil, err
	}
	if err := p.validator.ValidateSchemaString(schemaString); err != nil {
		return nil, err
	}

	challengeID, _ := strconv.ParseUint(challengeIDString, 10, 64)
	currTimestamp := uint64(time.Now().Unix())

	return &addProgressParams{
		userID:        msg.UserID,
		challengeID:   challengeID,
		schema:        []byte(schemaString),
		dateTimestamp: currTimestamp,
	}, nil
}

func (p parser) ParseListProgress(ctx context.Context, msg base.MsgData) (*listProgressParams, error) {
	tokens := base.GetTokens(msg.Msg)
	if !base.IsCorrectNumTokens(tokens, listNumTokens) {
		return nil, base.ErrInvalidTokenCount
	}

	challengeIDString := tokens[1]
	if err := p.validator.ValidateID(challengeIDString); err != nil {
		return nil, err
	}
	challengeID, _ := strconv.ParseUint(challengeIDString, 10, 64)

	return &listProgressParams{
		challengeID: challengeID,
		userID:      msg.UserID,
	}, nil
}

func NewParser(validator validator.Validator) Parser {
	return &parser{validator: validator}
}
