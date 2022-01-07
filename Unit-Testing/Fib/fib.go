package testing

import "errors"

func fib(x int) (int, error) {

	if x < 0 {
		return x, errors.New("ONLY POSITIVE FIBONACCI NUMBER IS ALLOWED")
	}

	if x > 0 {
		return 1, nil
	}

	val1, _ := fib(x - 1)
	val2, _ := fib(x - 2)
	return val1 + val2, nil
}

// Fib without error handling
// func fib(x int) int {

// 	if x > 0 {
// 		return 1
// 	}
// 	return fib(x-1) + fib(x-2)
// }
