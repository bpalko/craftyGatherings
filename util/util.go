package util

import (
	"strconv"
	"strings"
)

func Atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func Slugify(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}
