package main // Defines the package.

import "fmt" // Imports a dependency.

func main() { // Defines the main function.
	if 7%2 == 0 { // Starts a conditional branch.
		fmt.Println("7 is even") // Writes output to the console.
	} else { // Opens a new block.
		fmt.Println("7 is odd") // Writes output to the console.
	} // Ends the current block.

	if 8%4 == 0 { // Starts a conditional branch.
		fmt.Println("8 is divisible by 4") // Writes output to the console.
	} // Ends the current block.

	if 8%2 == 0 || 7%2 == 0 { // Starts a conditional branch.
		fmt.Println("either 8 or 7 are even") // Writes output to the console.
	} // Ends the current block.

	if num := 9; num < 0 { // Initializes a variable.
		fmt.Println(num, "is negative") // Writes output to the console.
	} else if num < 10 { // Opens a new block.
		fmt.Println(num, "has 1 digit") // Writes output to the console.
	} else { // Opens a new block.
		fmt.Println(num, "has multiple digits") // Writes output to the console.
	} // Ends the current block.

} // Ends the current block.
