package fibonachi

import (
	"testing"
)

func check(t *testing.T, n, r int, e bool) {
	answer, err := Fib(n)
	if (err != nil) != e {
		t.Fatalf("expected error %v, actual %v", e, err)
	}
	if answer != r {
		t.Fatalf("expected answer %d, actual %d", r, answer)
	}
}

func TestFib1(t *testing.T) {
	check(t, 1, 1, false)
}

func TestFib2(t *testing.T) {
	check(t, 2, 1, false)
}

func TestFib3(t *testing.T) {
	check(t, 3, 2, false)
}

func TestFib10(t *testing.T) {
	check(t, 10, 55, false)
}

func TestFib0(t *testing.T) {
	check(t, 0, 0, true)
}
