package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fib_old := 0
	fib_new := 0	
	return func() int {
		if fib_new == 0 {
			fib_old = 0
			fib_new = 1
			return 0
		} else {
			temp := fib_new
			fib_new = fib_old + fib_new
			fib_old = temp
			return fib_old
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
