package mylib

import "fmt"

func Defer() {
	// prints 3 2 1 0 before surrounding function returns
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}

	result := f()
	fmt.Printf("result: %d\n", result)
}

// f returns 42
func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}
