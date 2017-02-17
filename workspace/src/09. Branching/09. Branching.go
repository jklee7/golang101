// Go has 2 types of branching statements: IF and SWITCH

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	if response := "Correct!"; int(inputInt) == 2 {
		fmt.Println(response)
	} else if response := "Close!"; inputInt > 0 && inputInt < 4 {
		fmt.Println(response)
	} else {
		fmt.Println("Wrong!")
	}
}
