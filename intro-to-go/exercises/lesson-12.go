// Write a program that does the following:
// Loop through the numbers 1 through 10
// Call a function that prints out if the number is even or odd (like in Lesson 11).
// Refactor your program to add a new function that takes an integer and returns a boolean
// signifying if the integer is even, use that function inside your printing function.

package main

import "fmt"

func isEven(i int) bool {
	return i%2 == 0
}

func main() {
	var m = make(map[int]bool)

	for n := 1; n <= 10; n++ {
		if isEven(n) {
			m[n] = true
		} else {
			m[n] = false
		}
	}

	for k, v := range m {
		if v == true {
			fmt.Printf("%d is even\n", k)
		} else {
			fmt.Printf("%d is odd\n", k)
		}
	}
}
