// Go has many constants already defined, for example `math.Pi`.
// Print out the perimeter distance around a circle that has a radius of 5 feet.
// The formula for perimeter distance is `2 * math.Pi * radius`.

package main

import (
	"fmt"
	"math"
)

func main() {
	var radius = 5.0
	var perimeter = 2 * math.Pi * radius
	fmt.Println(perimeter)
}
