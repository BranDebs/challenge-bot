package postgres

import (
	"strconv"
	"strings"
)

const (
	delim = ","
)

func serializeUint64(vals ...uint64) string {
	var builder strings.Builder

	for i, val := range vals {
		if i != 0 {
			builder.WriteString(delim)
		}
		ustr := strconv.FormatUint(val, 10)
		builder.WriteString(ustr)

	}

	return builder.String()
}

func deserializeToUint64(str string) []uint64 {
	toks := strings.Split(str, delim)

	vals := make([]uint64, len(toks))
	for i, tok := range toks {
		v, _ := strconv.ParseUint(tok, 10, 64)
		vals[i] = v
	}

	return vals
}
