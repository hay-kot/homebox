package reporting

import (
	"strconv"
	"strings"
)

func parseSeparatedString(s string, sep string) ([]string, error) {
	list := strings.Split(s, sep)

	csf := make([]string, 0, len(list))
	for _, s := range list {
		trimmed := strings.TrimSpace(s)
		if trimmed != "" {
			csf = append(csf, trimmed)
		}
	}

	return csf, nil
}

func parseFloat(s string) float64 {
	if s == "" {
		return 0
	}
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func parseBool(s string) bool {
	b, _ := strconv.ParseBool(s)
	return b
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
