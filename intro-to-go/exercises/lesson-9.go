// Write a program that does the following:
// Creates an empty slice of integers and prints it out -> []
// Loops over the numbers 1 through 10 and appends them to the slice
// Prints out the slice -> [1 2 3 4 5 6 7 8 9 10]
// Prints out the first 5 numbers in the slice -> [1 2 3 4 5]

package main

import "fmt"

func main() {
	var s []int
	fmt.Println("Empty slice:", s)

	for n := 1; n <= 10; n++ {
		s = append(s, n)
	}

	fmt.Println("Filled slice:", s)

	newSlice := s[0:5]
	fmt.Println("New slice:", newSlice)
}
