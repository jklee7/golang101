// Go has 2 types of branching statements: IF and SWITCH

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Switches can also be used to determine type of variables.
// In this example, we take in an interface (i.e. void pointer)
// and write out the type.
func DetectType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case float32:
		fmt.Println("Float32")
	default:
		fmt.Println("Unknown")
	}
}

func main() {
	// Read from commandline using the bufio library.
	// We create a Reader object...
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What's 1 + 1?")

	// and then call the ReadString method
	input, _ := reader.ReadString('\n')

	// TrimSpace returns a slice of the string s
	// with all leading and trailing white space removed
	input = strings.TrimSpace(input)

	// Convert the input into an integer.
	// Invalid text (non-numbers) will be converted to '0'
	var inputInt int
	inputInt, _ = strconv.Atoi(input)

	// If statements in go follow the usual syntax:
	// if [optional statement]; [condition] {
	// } else if [condition] {
	// } else {
	// }
	//
	// Optional statements are only executed when the if statement
	// is executed.
	fmt.Println("Using IF/ELSE IF/ELSE")
	if response := "Correct!"; int(inputInt) == 2 {
		fmt.Println(response)
	} else if response := "Close!"; inputInt > 0 && inputInt < 4 {
		fmt.Println(response)
	} else {
		fmt.Println("Wrong!")
	}

	// The same thing above, using a SWITCH statement
	fmt.Println("\nUsing SWITCH")
	switch inputInt {
	case 2:
		fmt.Println("Correct!")
	case 1, 3:
		fmt.Println("Close!")
	default:
		fmt.Println("Wrong!")
	}

	// The same thing above, using an EMPTY SWITCH statement
	fmt.Println("\nUsing EMPTY SWITCH")
	switch {
	case inputInt == 2:
		fmt.Println("Correct!")
	case inputInt > 0 && inputInt < 4:
		fmt.Println("Close!")
	default:
		fmt.Println("Wrong!")
	}

	fmt.Println("\nDetecting variable types")
	DetectType("xyz123")
	DetectType(55)
	DetectType(105.4)
}
