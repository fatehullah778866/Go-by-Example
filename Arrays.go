package main // Defines the package.

import "fmt" // Imports a dependency.

func main() { // Defines the main function.
	var a [5]int // Declares a variable.
	fmt.Println("emp:", a) // Writes output to the console.

	a[4] = 100 // Executes a statement.
	fmt.Println("set:", a) // Writes output to the console.
	fmt.Println("get:", a[4]) // Writes output to the console.
	fmt.Println("len:", len(a)) // Writes output to the console.

	b := [5]int{1, 2, 3, 4, 5} // Initializes a variable.
	fmt.Println("dcl:", b) // Writes output to the console.

	b = [...]int{1, 2, 3, 4, 5} // Executes a statement.
	fmt.Println("dcl:", b) // Writes output to the console.

	b = [...]int{100, 3: 400, 500} // Executes a statement.
	fmt.Println("idx:", b) // Writes output to the console.
	var twoD [2][3]int // Declares a variable.
	for i := range 2 { // Initializes a variable.
		for j := range 3 { // Initializes a variable.
			twoD[i][j] = i + j // Executes a statement.
		} // Ends the current block.
	} // Ends the current block.
	fmt.Println("2d: ", twoD) // Writes output to the console.
	twoD = [2][3]int{ // Opens a new block.
		{1, 2, 3}, // Executes a statement.
		{1, 2, 3}, // Executes a statement.
	} // Ends the current block.
	fmt.Println("2d: ", twoD) // Writes output to the console.
} // Ends the current block.
