// Package swozibble provides a small consumer package for zibble.
package swozibble

import "github.com/xuancao/zibble"

// Process reverses input and formats it as uppercase output.
func Process(input string) string {
	return zibble.Shout(zibble.Reverse(input))
}
