package validator

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Validator interface {
	ValidateEndDateString(startTimestamp uint64, endTimestamp uint64) error
	ValidateSchemaString(schemaString string) error
	ValidateID(id string) error
}

var (
	ErrInvalidEndTimestamp = errors.New("end date cannot be earlier than current date")
	ErrInvalidJsonFormat   = errors.New("schema provided is not in json format")
	ErrInvalidIDFormat     = errors.New("invalid ID format, ID should be a positive integer")
)

type validator struct {
}

func (v validator) ValidateEndDateString(startTimestamp uint64, endTimestamp uint64) error {
	if endTimestamp <= startTimestamp {
		return ErrInvalidEndTimestamp
	}
	return nil
}

// TODO Deb
func (v validator) ValidateSchemaString(schemaString string) error {
	if !v.isJSON(schemaString) {
		return ErrInvalidJsonFormat
	}
	return nil
}

func (v validator) isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func (v validator) ValidateID(id string) error {
	_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return ErrInvalidIDFormat
	}
	return nil
}

func NewValidator() Validator {
	return &validator{}
}
