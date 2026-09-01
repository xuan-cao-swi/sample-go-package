// Command basic demonstrates the zibble package's functions.
package main

import (
	"fmt"

	zibble "github.com/xuan-cao-swi/sample-go-package"
)

func main() {
	fmt.Println(zibble.Reverse("hello"))
	fmt.Println(zibble.IsPalindrome("racecar"))
	fmt.Println(zibble.Shout("hi"))
	fmt.Println(zibble.NewID())
}
