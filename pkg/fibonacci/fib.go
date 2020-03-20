package fibonacci

import (
	"fmt"
)

func Fib(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("illegal number")
	}
	if n == 1 || n == 2 {
		return 1, nil
	}
	a, err := Fib(n - 1)
	if err != nil {
		return 0, err
	}
	b, err := Fib(n - 2)
	if err != nil {
		return 0, err
	}
	return a + b, nil
}
