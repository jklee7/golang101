package main

import "fmt"

func main() {
	// Arrays in Go are pretty much like in every other languages
	// - Have a fixed size
	// - Are Zero-valued at declaration (no need to initialize)
	// - Are NOT pointers
	// - Are a unique type themselves: combination of size and type
	var array [5]int

	for i := 0; i < 5; i++ {
		array[i] = i * 10
	}
	for i := range array {
		fmt.Println(array[i])
	}

	// Slices are similar to arrays, but are more flexible
	// - Are effectively pointers to arrays
	// - Have a fixed size, but can be reallocated
	// - Initialize, otherwise is NIL

	// Syntax:
	// Declaration: var sliceName[] type
	// Initialization: sliceName = make([]type, initialSize, maximumSize(optional) )
	var slice []int
	slice = make([]int, 3)
	slice[0] = 11
	slice[1] = 22
	slice[2] = 33

	for i := range slice {
		fmt.Println(slice[i])
	}

	// Shorthand method:
	slice2 := []int{100, 200, 300, 400, 500}

	for i := range slice2 {
		fmt.Println(slice2[i])
	}

	fmt.Println()
	// Slices can be further 'sliced up'.
	// Remember - a slice is just a pointer
	// slice = slice[startingIndex:endingIndex]
	// Note: does not include the ending index!
	sliced := slice2[1:3]
	for i := range sliced {
		fmt.Println(sliced[i])
	}

	fmt.Println()
	// You can also do a form of shorthand slicing:
	// by excluding the ending index, it automatically means the last index
	sliced = slice2[3:]
	for i := range sliced {
		fmt.Println(sliced[i])
	}

	fmt.Println()
	// Similarly by excluding the start index,
	// it automatically means the first index
	sliced = slice2[:3]
	for i := range sliced {
		fmt.Println(sliced[i])
	}

}
