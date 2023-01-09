package configs

import (
	"regexp"
	"strings"
)

// Convert Camel case string to normal case
func ToNormalCase(str string) string {

	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1} ${2}")
	return strings.ToLower(snake)
}

// Convert Integer to Float64
func IntToFloat64(value int) (newValue float64) {
	return float64(value)
}

// Convert Float64 to Integer
func Float64ToInt(value float64) (newValue int) {
	return int(value)
}
