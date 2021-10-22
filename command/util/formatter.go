package util

import "strings"

func CleanMarkdownMsg(msg string) string {
	specialChars := []string{"_", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
	for _, specialChar := range specialChars {
		msg = strings.ReplaceAll(msg, specialChar, "\\"+specialChar)
	}
	return msg
}
