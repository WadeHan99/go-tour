package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {
	z := float64(2)
	s := float64(0)
	for i := 0; i < 10; i++ {
		z = z - (z*z - x) / (2 * z)
		if math.Abs(z - s) < 1e-10 {
			break;
		}
		s = z
	}
	return z
}

func Sqrt2(x float64) float64 {
	var z float64 = 1.00
	for i := 0; i < 10; i++ {
		z = z - (z*z -x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println("Sqrt1     = ", Sqrt1(2))
	fmt.Println("Sqrt2     = ", Sqrt2(2))
	fmt.Println("math.Sqrt = ", math.Sqrt(2))
}

