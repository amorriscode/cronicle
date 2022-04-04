package utils

import (
	"io/fs"
	"strconv"
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

func GetIdFromArg(args []string, files []fs.FileInfo) int {
	n, err := strconv.Atoi(args[0])
	if err != nil || n == 0 || n > len(files) {
		return -1
	}
	return n - 1
}
