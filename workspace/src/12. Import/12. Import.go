// An example of how to import your own Go libraries
//
// In this example, I've created a library under ./exampleLibrary/exampleLibrary.go
// with a single function - PrintString(string).
// I then import the library and execute the function.

package main

import "./exampleLibrary"

func main() {
	exampleLibrary.PrintString("Example of importing your own GO libraries")
}
