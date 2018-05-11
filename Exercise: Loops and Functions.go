package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 0.5
	count := 0
	for i:=0; i<10; i++ {
		if z == (z*z - x) / (2*z) {
			break
		} else {
			z -= (z*z - x) / (2*z)
			count++
		}
	}
	fmt.Println(count)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

