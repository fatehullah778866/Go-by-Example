// package main identifies this file as an executable program
package main

import (
	"errors" // Used for the errors.As function
	"fmt"    // Used for formatting and printing
)

// Define a custom struct to hold error details.
// It keeps the 'arg' that caused the problem and a 'message'.
type argError struct {
	arg     int
	message string
}

// By adding this Error() method, the 'argError' struct now
// follows the "error" interface rules in Go.
func (e *argError) Error() string {
	// Return a formatted string combining the number and the message
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// f is a function that returns a result and an error.
func f(arg int) (int, error) {
	if arg == 42 {
		// Return our custom struct as the error.
		// We use & to return a pointer because our Error() method uses a pointer receiver.
		return -1, &argError{arg, "can't work with it"}
	}
	// Return the result and nil if no error occurs
	return arg + 3, nil
}

func main() {

	// Call the function with 42 to trigger the error
	_, err := f(42)

	// Create a variable 'ae' of type pointer to argError.
	// We will use this to "catch" the specific error data.
	var ae *argError

	// errors.As checks if 'err' is actually an '*argError'.
	// If it is, it saves the data into 'ae' and returns true.
	if errors.As(err, &ae) {
		// Now we can access the custom fields directly!
		fmt.Println(ae.arg)     // Prints: 42
		fmt.Println(ae.message) // Prints: can't work with it
	} else {
		// This runs if the error was a different type
		fmt.Println("err doesn't match argError")
	}
}
