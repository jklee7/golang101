// Go allows for passing functions as parameters
// In this example, we can pass more specialised functions
// (e.g. Basic arithmetic) into a more general function

package main

import "fmt"

// Our general 'Math' function take 2 numbers and a function.
// Whatever function is passed will be executed on the 2 int parameters
func Math(x, y int, do func(int, int) int) int {
	return do(x, y)
}

// Our more specialised functions - Add, Subtract, Multiply, Divide
func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	return x * y
}

func Divide(x, y int) int {
	return x / y
}

// We can also declare a function as its own type
type DoMath func(int, int) int

// This simplifies passing it as a parameter
func Math2(x, y int, do DoMath) int {
	return do(x, y)
}

// You can even specify functions as a return value...
// i.e. a function that returns functions

func CreateFunction(custom int) DoMath {
	return func(x, y int) int {
		return x + y + custom
	}
}

func main() {
	// We can then pass the function we want to execute on the 2 integers
	// as a parameter to the more general Math function
	fmt.Println(Math(20, 10, Add))
	fmt.Println(Math(20, 10, Subtract))
	fmt.Println(Math(20, 10, Multiply))
	fmt.Println(Math(20, 10, Divide))

	// Using functions which take function types
	fmt.Println(Math2(20, 10, Add))
	fmt.Println(Math2(20, 10, Subtract))
	fmt.Println(Math2(20, 10, Multiply))
	fmt.Println(Math2(20, 10, Divide))

	// Using the CreateFunction function to generate a function!
	fmt.Println(Math2(20, 10, CreateFunction(100)))
}
