package utils

import (
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile(`(.)[A-Z][a-z]+`)
	matchAllCap   = regexp.MustCompile(`([a-z0-9])([A-Z])`)
)
func CamelToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}
func NormailizeString(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}