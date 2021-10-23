package base

import (
	"errors"
	"regexp"
	"strings"
)

const (
	regexExpr = `[^\s']+|'([^']*)'`
)

var (
	ErrInvalidTokenCount = errors.New("number of tokens provided is invalid")
)

func IsCorrectNumTokens(tokens []string, numTokens int) bool {
	return len(tokens) == numTokens
}

func GetTokens(msg string) []string {
	regExp := regexp.MustCompile(regexExpr)
	rawTokens := regExp.FindAllString(msg, -1)

	// remove ' string around tokens
	for i, token := range rawTokens {
		rawTokens[i] = strings.ReplaceAll(token, "'", "")
	}

	return rawTokens
}
