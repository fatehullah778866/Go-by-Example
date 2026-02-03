package main // Defines the package.

import "fmt" // Imports a dependency.

func sum(nums ...int) { // Defines the sum function.
	fmt.Print(nums, " ") // Writes output to the console.
	total := 0 // Initializes a variable.
	for _, num := range nums { // Initializes a variable.
		total += num // Executes a statement.
	} // Ends the current block.
	fmt.Println(total) // Writes output to the console.
} // Ends the current block.

func main() { // Defines the main function.
	sum(1, 2) // Executes a statement.
	sum(1, 2, 3) // Executes a statement.
	nums := []int{1, 2, 3, 4} // Initializes a variable.
	sum(nums...) // Executes a statement.
} // Ends the current block.
