package main // Defines the package.

import ( // Starts the import block.
	"fmt" // Executes a statement.
	"slices" // Executes a statement.
) // Ends the import block.

func main() { // Defines the main function.
	var s []string // Declares a variable.
	fmt.Println("uninit:", s, s == nil, len(s) == 0) // Writes output to the console.

	s = make([]string, 3) // Allocates a new data structure.
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // Writes output to the console.
	s[0] = "a" // Executes a statement.
	s[1] = "b" // Executes a statement.
	s[2] = "c" // Executes a statement.
	fmt.Println("set:", s) // Writes output to the console.
	fmt.Println("get:", s[2]) // Writes output to the console.
	fmt.Println("len:", len(s)) // Writes output to the console.

	s = append(s, "d") // Appends to a slice.
	s = append(s, "e", "f") // Appends to a slice.
	fmt.Println("apd:", s) // Writes output to the console.

	c := make([]string, len(s)) // Initializes a variable.
	copy(c, s) // Copies slice elements.
	fmt.Println("cpy:", c) // Writes output to the console.

	l := s[2:5] // Initializes a variable.
	fmt.Println("sl1:", l) // Writes output to the console.

	l = s[:5] // Executes a statement.
	fmt.Println("sl2:", l) // Writes output to the console.

	l = s[2:] // Executes a statement.
	fmt.Println("sl3:", l) // Writes output to the console.

	t := []string{"g", "h", "i"} // Initializes a variable.
	fmt.Println("dcl:", t) // Writes output to the console.

	t2 := []string{"g", "h", "i"} // Initializes a variable.
	if slices.Equal(t, t2) { // Starts a conditional branch.
		fmt.Println("t == t2") // Writes output to the console.
	} // Ends the current block.

	twoD := make([][]int, 3) // Initializes a variable.
	for i := range 3 { // Initializes a variable.
		innerLen := i + 1 // Initializes a variable.
		twoD[i] = make([]int, innerLen) // Allocates a new data structure.
		for j := range innerLen { // Initializes a variable.
			twoD[i][j] = i + j // Executes a statement.
		} // Ends the current block.
	} // Ends the current block.
	fmt.Println("2d: ", twoD) // Writes output to the console.
} // Ends the current block.
