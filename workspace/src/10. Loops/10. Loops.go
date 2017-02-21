// Go has only 1 type of loop structure: FOR
//
// However, FOR loops in go are flexible, and can
// be used the same way as traditional WHILE loops too

package main

import "fmt"

func main() {

	//'Traditional' FOR loop
	fmt.Println("Traditional FOR loop")
	fmt.Println("====================")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//'While' loop
	fmt.Println("FOR used like a WHLIE loop")
	fmt.Println("==========================")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//Infinite loop
	fmt.Println("FOR infinite loops")
	fmt.Println("==================")
	i = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++

	}

	//Using 'continue' to short-circuit loops
	//In the below example, continue is used to skip printing
	//odd numbers
	fmt.Println("Using 'continue' to shortcircuit loops")
	fmt.Println("======================================")
	i = 0
	for {
		if i >= 10 {
			break
		}

		// 'continue' skips all the code below and restarts at the top of the loop
		if i%2 != 0 {
			i++
			continue
		}

		fmt.Println(i)
		i++
	}

	//Looping through Ranges
	//Slices are like Arrays. Unlike arrays though,
	//slices are flexible, and have no fixed length
	fmt.Println("Looping through ranges")
	fmt.Println("======================")
	slice := []int{
		1, 2, 3, 4, 5,
	}

	//When looping through ranges, the syntax is:
	//for [index] [value] := range [slice]
	for i, v := range slice {
		fmt.Println("Index ", i, ": ", v)
	}

}
