package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	t := 1.0
	c := 0
	for {
		c++
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-t) < 0.000000000000001 {
			return z, c
		}
		t = z
	}
}

func main() {
	i := 1.0
	for i <= 10 {
		value, count := Sqrt(i)
		fmt.Println(i, " -> ", value, ", ", count, "回")
		i++
	}
	value, count := Sqrt(9)
	fmt.Println("9 -> ", value, ", ", count, "回")
}
