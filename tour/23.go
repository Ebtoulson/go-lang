// As a simple way to play with functions and loops,
// implement the square root function using Newton's method.

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var y, z float64 = 0, 1
	for {
		z = z - (math.Pow(z, 2)-x)/(2*z)
		if math.Abs(z-y) < 1e-15 {
			break
		}
		y = z
	}
	return y
}

func main() {
	fmt.Println(Sqrt(25))
	fmt.Println(math.Sqrt(25))
}
