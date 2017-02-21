// An example of creating your own Go library.
// Note: We don't have a 'main' function because this is a library, not an actual GO
// program.

package exampleLibrary

import "fmt"

// Note: the CAPITAL first letter denotes a 'public' or 'exported' function.
// Functions named with a lowercase letter are considered private.
func PrintString(example string) {
	fmt.Println(example)
}
