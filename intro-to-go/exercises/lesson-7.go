// Combine a for loop with switch to count to 10 and
// print whether the number is less than 5, equal to 5, or greater than 5.

package main

import "fmt"

func main() {
	for s := 1; s <= 10; s++ {
		switch {
		case s < 5:
			fmt.Printf("The number %d is less than 5\n", s)
		case s == 5:
			fmt.Printf("The number %d is equal to 5\n", s)
		case s > 5:
			fmt.Printf("The number %d is greater than 5\n", s)
		}
	}

}
