package base

import (
	"encoding/json"
	"strings"
	"time"
)

const (
	dateLayout = "02-01-2006 15:04:00"
)

func CleanMarkdownMsg(msg string) string {
	specialChars := []string{"_", "[", "]", "(", ")", "~", "`", ">", "#", "+", "-", "=", "|", "{", "}", ".", "!"}
	for _, specialChar := range specialChars {
		msg = strings.ReplaceAll(msg, specialChar, "\\"+specialChar)
	}
	return msg
}

func FormatTimestampToDate(timestamp int64) string {
	convertedTime := time.Unix(timestamp, 0)
	singapore, err := time.LoadLocation("Asia / Singapore")
	if err != nil {
		return convertedTime.Format(dateLayout)
	}

	singaporeTime := convertedTime.In(singapore)
	return singaporeTime.Format(dateLayout)
}

func FormatSchemaValue(value []byte) map[string]interface{} {
	var res map[string]interface{}
	if err := json.Unmarshal(value, &res); err != nil {
		return nil
	}
	return res
}
