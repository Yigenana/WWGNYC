// Write a for loop that only executes once.
// Then write a for loop with no condition statement
// that exits when the count is equal to 100.

package main

import "fmt"

func main() {
	for s := 1; s < 2; s++ {
		fmt.Println("One loop")
	}

	var t = 1
	for t < 101 {
		fmt.Println("Iteration number", t)
		t++
	}

}
