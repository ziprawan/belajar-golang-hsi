package utils

import "strconv"

// Parse string into unsigned integer
// If the parse failed, return 0
func SafeParseUint(str string) uint {
	parsed, _ := strconv.ParseUint(str, 10, 0)
	return uint(parsed)
}
