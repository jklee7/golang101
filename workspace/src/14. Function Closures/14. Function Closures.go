// Because Go implements first-class functions, functions can be
// stored in variables.
// This brings the concept of closures - basically, variables declared
// inside a function are persistent.
//
// A typical demonstration of this is printing the fibonacci sequence.

package main

import "fmt"

func fib() func() int {
	x, y := 0, 1

	// x & y are declared within the function.
	// If you just call the function on its own, you'd
	// only ever get the value '1', because the values
	// are discarded after the function finishes
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	// When assigned to a variable though,
	// the variables declared within the function persist.
	f := fib()
	// When assigned to a variable, the variable can be called like
	// a function, but it retains any values declared in the
	// original function.
	fmt.Println(f(), f(), f(), f(), f(), f())

	g := fib()
	fmt.Println(g(), g(), g(), g(), g())
}
