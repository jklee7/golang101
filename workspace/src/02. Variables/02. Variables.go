package main

import "fmt"

func main() {
	// Variables in Go are declared in the following syntax:
	//	var variableName type
	var message string = "Hello World!"

	// Multiple variables can be declared in 1 line
	var a, b, c int = 10, 20, 30

	// You can also use the := shortcut method of initialising variables,
	// but only when declaring variables within a function
	shortcut := "shortcut"
	x, y, z := 1, 2, 3

	fmt.Println(message, a, b, c, shortcut, x, y, z)
}
