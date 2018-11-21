package printem

import (
	"fmt"
)

// Printem takes a variable number of integer arguments and prints them on separate lines
func Printem(args ...int) {
	for _, i := range args {
		fmt.Println(i)
	}
}
