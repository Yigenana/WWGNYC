// Combine a For loop with an If/Else to count to 10
// and print out if a number is even or odd.

package main

import "fmt"

func main() {
	for s := 1; s <= 10; s++ {
		if s%2 == 0 {
			fmt.Printf("The number %d is even\n", s)
		} else {
			fmt.Printf("The number %d is odd\n", s)
		}
	}

}
