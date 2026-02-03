package main // Defines the package.

import "fmt" // Imports a dependency.

func intSeq() func() int { // Defines the intSeq function.
	i := 0 // Initializes a variable.
	return func() int { // Returns a value from the function.
		i++ // Executes a statement.
		return i // Returns a value from the function.
	} // Ends the current block.
} // Ends the current block.

func main() { // Defines the main function.

	nextInt := intSeq() // Initializes a variable.

	fmt.Println(nextInt()) // Writes output to the console.
	fmt.Println(nextInt()) // Writes output to the console.
	fmt.Println(nextInt()) // Writes output to the console.

	newInts := intSeq() // Initializes a variable.
	fmt.Println(newInts()) // Writes output to the console.
} // Ends the current block.
