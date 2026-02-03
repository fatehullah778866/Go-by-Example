package main // Defines the package.

// import "fmt"
// import "math"

import ( // Starts the import block.
	"fmt" // Executes a statement.
	"math" // Executes a statement.
) // Ends the import block.

const s string = "constant" // Declares a constant.

func main() { // Defines the main function.
	fmt.Println(s) // Writes output to the console.
	const n = 500000000 // Declares a constant.
	const d = 3e20 / n // Declares a constant.
	fmt.Println(d) // Writes output to the console.

	fmt.Println(int64(d)) // Writes output to the console.

	fmt.Println(math.Sin(n)) // Writes output to the console.
} // Ends the current block.
