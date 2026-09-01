// Command basic demonstrates the zibble package's functions.
package main

import (
	"fmt"

	"github.com/xuancao/zibble"
)

func main() {
	fmt.Println(zibble.Reverse("hello"))
	fmt.Println(zibble.IsPalindrome("racecar"))
	fmt.Println(zibble.Shout("hi"))
}
