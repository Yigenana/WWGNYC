// Write a program that does the following:
// Creates a map from integers to booleans
// Loops through the numbers 1 through 10 and sets the map where the key is the number
// and the value is whether the number is even (HINT: use an If/Else or Switch!)
// Prints out the map -> map[1:false 2:true 3:false 4:true 5:false 6:true 7:false 8:true 9:false 10:true]
// Prints out whether 4 is even using variables -> The number 4 is even: true

package main

import "fmt"

func main() {
	var m = make(map[int]bool)

	for n := 1; n <= 10; n++ {
		if n%2 == 0 {
			m[n] = true
		} else {
			m[n] = false
		}
	}

	fmt.Println("Filled map:", m)

	var mapValue = m[4]
	fmt.Println("The number 4 is even:", mapValue)
}
