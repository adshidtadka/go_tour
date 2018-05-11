package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 0.5
	for i := 0; i < 10; i++ {
		if z == (z*z-x)/(2*z) {
			break
		} else {
			z -= (z*z - x) / (2 * z)
		}
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	if val, err := Sqrt(2); err == nil {
		fmt.Println(val)
	} else {
		fmt.Println(err)
	}
	if val, err := Sqrt(-2); err == nil {
		fmt.Println(val)
	} else {
		fmt.Println(err)
	}
	fmt.Println(Sqrt(-2))
}
