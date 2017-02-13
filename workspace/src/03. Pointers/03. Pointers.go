package main

import "fmt"

func main() {
	// standard string variable
	var message string = "Hello World"

	// pointers are denoted by *
	// memory addresses of a variable are denoted with &
	var greeting *string = &message

	// we can reference the value of what the pointer is pointing to using *
	fmt.Println(message, greeting, *greeting)

	// can also change the value of what the pointer is referencing using *
	*greeting = "Goodbye World"
	fmt.Println(message, greeting, *greeting)
}
