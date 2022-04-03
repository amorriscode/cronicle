package utils

import (
	"strings"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BoolPtr(b bool) *bool {
	return &b
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func TruncateText(s string, max int) string {
	t := strings.TrimSpace(s)

	if max > len(t) {
		return t
	}

	if strings.LastIndex(s[:max], " ") == -1 {
		return s[:max]
	}

	return s[:strings.LastIndex(s[:max], " ")] + "..."
}
