package main // Defines the package.

import "fmt" // Imports a dependency.

func fact(n int) int { // Defines the fact function.
	if n == 0 { // Starts a conditional branch.
		return 1 // Returns a value from the function.
	} // Ends the current block.
	return n * fact(n-1) // Returns a value from the function.
} // Ends the current block.

func main() { // Defines the main function.
	fmt.Println(fact(7)) // Writes output to the console.

	var fib func(n int) int // Declares a variable.

	fib = func(n int) int { // Opens a new block.
		if n < 2 { // Starts a conditional branch.
			return n // Returns a value from the function.
		} // Ends the current block.

		return fib(n-1) + fib(n-2) // Returns a value from the function.
	} // Ends the current block.

	fmt.Println(fib(7)) // Writes output to the console.
} // Ends the current block.
