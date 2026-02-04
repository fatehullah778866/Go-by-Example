// package main identifies this file as an executable program
package main

import (
	"errors" // Package used to create and manipulate error values
	"fmt"    // Package used for printing and formatting
)

// f is a function that returns a result (int) and an error (error).
// This is the standard "Go way" to handle potential problems.
func f(arg int) (int, error) {
	// Check for a specific condition that causes a failure
	if arg == 42 {
		// Return -1 and a new error message
		return -1, errors.New("can't work with 42")
	}

	// If everything is okay, return the result and 'nil' for the error
	return arg + 3, nil
}

// We can define "sentinel errors" as global variables.
// These are specific error values we can check for later.
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		// Return the specific "Out of Tea" error
		return ErrOutOfTea
	} else if arg == 4 {
		// Use %w to "wrap" an error. This adds context but keeps the original error inside.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil // No error occurred
}

func main() {
	// Loop through a slice of numbers to test function f
	for _, i := range []int{7, 42} {

		// Check the result 'r' and the error 'e' at the same time
		if r, e := f(i); e != nil {
			// If e is not nil, something went wrong
			fmt.Println("f failed:", e)
		} else {
			// If e is nil, the function was successful
			fmt.Println("f worked:", r)
		}
	}

	// A modern Go loop: range over an integer (0 to 4)
	for i := range 5 {
		if err := makeTea(i); err != nil {

			// errors.Is is the best way to check if an error is a specific type.
			// It works even if the error was "wrapped" with fmt.Errorf.
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				// If it's an error we didn't expect
				fmt.Printf("unknown error: %s\n", err)
			}
			continue // Skip to the next iteration of the loop
		}

		// If err was nil, the tea was made successfully
		fmt.Println("Tea is ready!")
	}
}
