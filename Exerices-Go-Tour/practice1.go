package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) (z float64) {
    z = x 
    i := 0
    oldz := 0.0
    //for i < 10 {
    for math.Abs(oldz - z) > 0 {
        if i > 100 {
            fmt.Println("It repeated over 100 times")
            return
        }
        oldz = z
        z = z - (z * z - x) / (2 * z)
        i += 1
    }
    
    fmt.Println("How many times repeated:", i)
    return 
}

func main() {
    var testcase float64 = 3
    fmt.Println("Square root by Newton's method:", Sqrt(testcase))
    fmt.Println("Square root by math package:", math.Sqrt(testcase))
}
