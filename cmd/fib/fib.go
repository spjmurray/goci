package main

import (
	"fmt"
	"flag"

	"github.com/spjmurray/goci/pkg/fibonacci"
)

func main() {
	n := flag.Int("n", 1, "fibonacci number")
	flag.Parse()

	v, err := fibonacci.Fib(*n)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("fibonacci", *n, "=", v)
}
