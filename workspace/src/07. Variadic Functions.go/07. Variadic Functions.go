// Variadic functions are functions which support
// a variable number of parameters of a certain type.

package main

import "fmt"

// Variadics are denoted with ...type
// This effectively turns the parameter into a slice (dynamic arrays in Go)
func PrintStrings(varString ...string) {
	fmt.Println("Number of parameters: ", len(varString))
	for i := 0; i < len(varString); i++ {
		fmt.Println(varString[i])
	}
}

func main() {
	PrintStrings("Hello", "World", "this", "is", "the", "end")
}
