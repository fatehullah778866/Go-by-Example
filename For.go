package main // Defines the package.

import "fmt" // Imports a dependency.

func main() { // Defines the main function.
	i := 1 // Initializes a variable.
	for i <= 3 { // Starts a loop.
		fmt.Println(i) // Writes output to the console.
		i = i + 1 // Executes a statement.
	} // Ends the current block.
	for j := 0; j < 3; j++ { // Initializes a variable.
		fmt.Println(j) // Writes output to the console.
	} // Ends the current block.

	for i := range 3 { // Initializes a variable.
		fmt.Println("range", i) // Writes output to the console.
	} // Ends the current block.

	for { // Starts a loop.
		fmt.Println("loop") // Writes output to the console.
		break // Exits the current block.
	} // Ends the current block.

	for n := range 6 { // Initializes a variable.
		if n%2 == 0 { // Starts a conditional branch.
			continue // Continues to the next loop iteration.
		} // Ends the current block.
		fmt.Println(n) // Writes output to the console.
	} // Ends the current block.

} // Ends the current block.
