package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, zn := x/2, x/2

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)

		if math.Abs(zn-z) < 0.000001 {
			fmt.Println(i)
			return z
		}

		zn = z
	}

	return z
}

func main() {
	x := 31.0
	fmt.Println(Sqrt(x))
	fmt.Println(math.Sqrt(x))
}
