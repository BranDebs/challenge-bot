package validator

type Validator interface {
	ValidateEndDateString(endDateString string) error
	ValidateSchemaString(schemaString string) error
	ValidateChallengeID(challengeID string) error
}

type validator struct {
}

func (v validator) ValidateEndDateString(endDateString string) error {
	panic("implement me")
}

func (v validator) ValidateSchemaString(schemaString string) error {
	panic("implement me")
}

func (v validator) ValidateChallengeID(challengeID string) error {
	panic("implement me")
}

func NewValidator() Validator {
	return &validator{}
}
