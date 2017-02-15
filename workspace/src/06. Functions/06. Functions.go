package main

import "fmt"

// Functions are declared using the keyword: func
// func functionName (parameter parameterType)
func Greet(name string) {
	fmt.Println(name)
}

// If your function has a return value, specify
// the return type after the function parameter:
// func functionName (parameter parameterType) returnType
func Add(x, y int) int {
	return x + y
}

// In Go, functions can return multiple values.
// This is done by specifying the return types in brackets
func AddTwice(x, y int) (int, int) {
	return x + y, x + y
}

// Go also has the concept of named returns.
// Declare the variable name in the return portion of the function
func AddThrice(x, y int) (value1 int, value2 int, value3 int) {
	// assign the values to return variables
	value1 = x + y
	value2 = x + y
	value3 = x + y

	// return
	return
}

func main() {
	Greet("Hello World")
	fmt.Println(Add(1, 2))

	x, y := AddTwice(1, 2)
	fmt.Println(x, y)

	// If you don't want some of the return type,
	// you can 'throwaway' return values using _
	_, z := AddTwice(3, 4)
	fmt.Println(z)
	fmt.Println(AddThrice(2, 4))
}
