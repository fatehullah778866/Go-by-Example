package main // Defines the package.

import ( // Starts the import block.
	"fmt" // Executes a statement.
	"maps" // Executes a statement.
) // Ends the import block.

func main() { // Defines the main function.

	m := make(map[string]int) // Initializes a variable.

	m["k1"] = 7 // Executes a statement.
	m["k2"] = 13 // Executes a statement.

	fmt.Println("map:", m) // Writes output to the console.

	v1 := m["k1"] // Initializes a variable.
	fmt.Println("v1:", v1) // Writes output to the console.

	v3 := m["k3"] // Initializes a variable.
	fmt.Println("v3:", v3) // Writes output to the console.

	fmt.Println("len:", len(m)) // Writes output to the console.

	delete(m, "k2") // Executes a statement.
	fmt.Println("map:", m) // Writes output to the console.

	clear(m) // Executes a statement.
	fmt.Println("map:", m) // Writes output to the console.

	_, prs := m["k2"] // Initializes a variable.
	fmt.Println("prs:", prs) // Writes output to the console.

	n := map[string]int{"foo": 1, "bar": 2} // Initializes a variable.
	fmt.Println("map:", n) // Writes output to the console.

	n2 := map[string]int{"foo": 1, "bar": 2} // Initializes a variable.
	if maps.Equal(n, n2) { // Starts a conditional branch.
		fmt.Println("n == n2") // Writes output to the console.
	} // Ends the current block.
} // Ends the current block.
