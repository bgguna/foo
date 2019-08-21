package main

import (
	"fmt"
	"sort"
)

func main() {

	beatles := map[string]string{}

	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"

	// Loop through the map and print out the key and value
	for key, value := range beatles {
		fmt.Printf("%v \tplays the %v\n", key, value)
	}

	// Look up the key `Bob` and detect that it wasn't found by printing `not found`
	key := "Bob"
	if value, ok := beatles[key]; ok {
		fmt.Printf("Found member: %v who plays %v\n", key, value)
	} else {
		fmt.Printf("Not found\n")
	}

	// Declare a slice of strings called `keys` and collect all the keys from the map.
	keys := make([]string, 0, len(beatles))
	for k := range beatles {
		keys = append(keys, k)
	}

	// sort the keys
	sort.Strings(keys)

	// print out the keys
	fmt.Printf("%+v\n", keys)
}
