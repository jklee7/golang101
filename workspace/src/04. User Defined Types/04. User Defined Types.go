package main

import "fmt"

// User defined types with the following syntax:
// type typeName [variableType]
// In this example, we are created a new type called 'Title'
// (which is really just a string)
type Title string

// The more useful use of types applies to structs.
// In Go, structs are the way to create user-defined concrete types.
// The design of the struct type is unique compared to other programming languages.
//
// In simple terms, a struct is a collection of fields or properties.
// Unlike other object-oriented languages, Go does not provide a “class” keyword.
// Instead, in Go you’ll find structs — a lightweight version of classes.
//
// A struct type is exported into other packages if the name of the struct starts with a capital letter.

type Salutation struct { //Note the CAPITALIZATION, which effectively makes this Public!
	name     string //No need for public/private/protected
	greeting string
}

func main() {
	var example Title = "User Defined Type Example"
	fmt.Println(example)

	// Instantiate a user-defined struct type with the format
	// structType{}
	// *Note the use of curly brackets

	var x = Salutation{"John", "Hello"}
	fmt.Println(x.name, x.greeting)

	// Also possible to instantiate values in different order, by specifying the field name
	x = Salutation{greeting: "Hello", name: "John"}
	fmt.Println(x.name, x.greeting)

	// Can set fields directly too
	x.name = "Jack"
	x.greeting = "Goodbye"
	fmt.Println(x.name, x.greeting)
}
