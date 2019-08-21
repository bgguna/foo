package main

import "fmt"

func main() {
	parents := []string{"Carol", "Mike"}
	kids := []string{"Marcia", "Jan", "Cindy", "Greg", "Peter", "Bobby"}

	family := append(parents, kids...)
	fmt.Printf("Family: %v\n", family)

	fmt.Printf("Len of family: %v\n", len(family))
	fmt.Printf("Cap of family: %v\n", cap(family))

	youngest := family[len(family)-3:]
	fmt.Printf("Youngest: %v\n", youngest)

	extras := make([]string, 1, 1)
	extras[0] = "Alice"
	fmt.Println(extras)
}
