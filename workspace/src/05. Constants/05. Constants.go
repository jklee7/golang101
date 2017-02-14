package main

import "fmt"

// In computer programming, a constant is a value that cannot be
// altered by the program during normal execution,
// i.e. the value is constant.

// In Go, constants are declare with the keyword 'const'
// You can declare a single constant:
const Pi = 3.142

// Or multiple constants with a const block:
const (
	CCVisa       string = "Visa"
	CCMastercard string = "MasterCard"
	CCAmex       string = "American Express"
)

// Go has a special keyword: iota.
// It is a very general construct that is used as an
// automatically incrementing counter in const declaration blocks.
// With this automatically incrementing counter,
// it gives you a “short-hand” notation for assigning values to constants.

// With the start of every const block, iota gets reset to 0,
// so it can be used repeatedly.

const (
	a = iota //0
	b = iota //1
	c = iota //2
)

// One thing to note - in const blocks, you don't have to assign
// iota to every variable you want to be an iota.
// simply emit assigning a value in every following variable,
// and it will automatically be assigned 'iota'
const (
	w string = "blah"
	x        = iota //1
	y               //2
	z               //3
)

func main() {
	fmt.Println(Pi, CCVisa, CCMastercard, CCAmex)
	fmt.Println(a, b, c)
	fmt.Println(w, x, y, z)
}
