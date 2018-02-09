package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	if(x < 0) {
		return x, ErrNegativeSqrt(x)
	}
	z := x
    i := 0
    oldz := 0.0
    //for i < 10 {
    for math.Abs(oldz - z) > 0 {
        if i > 100 {
            fmt.Println("It repeated over 100 times")
            return z, nil
        }
        oldz = z
        z = z - (z * z - x) / (2 * z)
        i += 1
    }
    
    fmt.Println("How many times repeated:", i)
	return z, nil
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
