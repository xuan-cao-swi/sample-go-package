// Package zibble provides small string utility helpers.
package zibble

import (
	"strings"

	"github.com/google/uuid"
)

// Reverse returns s with its characters in reverse order.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome reports whether s reads the same forwards and backwards,
// ignoring case.
func IsPalindrome(s string) bool {
	lower := strings.ToLower(s)
	return lower == Reverse(lower)
}

// Shout uppercases s and appends an exclamation mark.
func Shout(s string) string {
	return strings.ToUpper(s) + "!"
}

// NewID returns a new random UUID string.
func NewID() string {
	return uuid.NewString()
}
