// Maps in GO are basically key-value pairs, similar to
// dictionaries in Python.

package main

import "fmt"

func main() {
	// Syntax: var [mapName] map [[key type]][value type]
	var websiteMap map[int]string

	// After declaring a map, you have to initialise it with 'make'
	websiteMap = make(map[int]string)

	// Add key:value pairs to the map
	websiteMap[10] = "Youtube"
	websiteMap[11] = "Facebook"
	websiteMap[12] = "Ebay"

	// Access the value by referencing the map key
	fmt.Println(websiteMap[10])
	fmt.Println(websiteMap[11])
	fmt.Println(websiteMap[12])

	// Maps can also be declared and initialised in 'shorthand'
	nameMap := map[int]string{
		1: "Walter",
		2: "Jessie",
		3: "Skyler",
		4: "Los Pollos Hermanos",
	}

	fmt.Println(nameMap[1])
	fmt.Println(nameMap[2])
	fmt.Println(nameMap[3])
	fmt.Println(nameMap[4])

	// Updating a map...
	nameMap[1] = "White"
	fmt.Println(nameMap[1])

	// Deleting from a map...
	delete(nameMap, 4)
	fmt.Println(nameMap[4])

	// Checking for existence in map
	if value, exists := nameMap[1]; exists {
		fmt.Println(value, "is in the Map!")
	}

}
