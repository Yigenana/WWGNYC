// Modify your program from lesson 10 to use a range to iterate the map to output the following:
// 1 is odd 2 is even 3 is odd 4 is even 5 is odd 6 is even 7 is odd 8 is even 9 is odd 10 is even
// HINT: use a for loop with range to get the key and value of each entry in the map,
// then use an If/Else on the value.

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

	for k, v := range m {
		if v == true {
			fmt.Printf("%d is even\n", k)
		} else {
			fmt.Printf("%d is odd\n", k)
		}
	}
}
