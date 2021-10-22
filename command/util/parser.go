package util

import (
	"regexp"
)

const (
	regexExpr            = `[^\s']+|'([^']*)'`
	InvalidTokenCountErr = "number of tokens provided is invalid"
)

func IsCorrectNumTokens(tokens []string, numTokens int) bool {
	return len(tokens) == numTokens
}

func GetTokens(msg string) []string {
	regExp := regexp.MustCompile(regexExpr)
	return regExp.FindAllString(msg, -1)
}
