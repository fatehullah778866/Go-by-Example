// package main identifies this file as an executable program
package main

// import "fmt" allows us to print text to the console
import "fmt"

// This function takes an integer 'ival' as an argument.
// In Go, this is "passed by value," meaning the function gets a copy of the number.
func zeroval(ival int) {
	// This changes the local copy to 0, but it does not affect the original variable
	ival = 0
}

// This function takes a pointer to an integer ('iptr').
// The *int means it stores the memory address of an integer.
func zeroptr(iptr *int) {
	// The * operator "dereferences" the pointer.
	// This means we go to the actual memory address and change the value stored there.
	*iptr = 0
}

func main() {
	// Initialize variable 'i' with the value 1
	i := 1
	// Print the starting value of i (which is 1)
	fmt.Println("initial:", i)

	// Call zeroval and pass 'i'.
	// The function receives a copy, so the original 'i' stays 1.
	zeroval(i)
	fmt.Println("zeroval:", i)

	// Call zeroptr and pass the memory address of 'i' using the & operator.
	// This allows the function to reach the original variable and change it.
	zeroptr(&i)
	// Now 'i' has been changed to 0
	fmt.Println("zeroptr:", i)

	// Print the memory address where 'i' is stored.
	// It will look like a hexadecimal code (e.g., 0xc0000120b0).
	fmt.Println("pointer:", &i)
}
