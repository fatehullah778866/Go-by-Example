package main

// What: Declares that this file belongs to the main package.
// Why: The "main" package is special in Go; it defines an executable program, not a library.
// When: Every standalone Go program needs a main package to run.
// Where: At the very top of the Go file.

import "fmt"

// What: Imports the "fmt" package, which provides formatting for input/output.
// How: Allows us to use functions like fmt.Println for printing to the console.
// Why: Without importing "fmt", you cannot print messages to the standard output.
// When: Needed whenever we want to output information or format text.

func main() {
	// What: Defines the main function, the entry point of a Go program.
	// How: The program starts executing from the main() function.
	// Why: Every executable Go program must have a main() function in package main.
	// When: Called automatically by the Go runtime when the program starts.

	messages := make(chan string)
	// What: Creates a new channel of type string and assigns it to the variable "messages".
	// How: The make() function allocates and initializes the channel.
	// Why: Channels are used in Go for communication between goroutines (concurrent functions).
	// When: Used here to send a message from one goroutine to another.
	// Where: In concurrent programming, channels are critical for safe data transfer between goroutines.

	go func() { messages <- "ping" }()
	// What: Launches an anonymous goroutine that sends the string "ping" into the "messages" channel.
	// How: The "go" keyword starts the function concurrently in a separate goroutine.
	// Why: To demonstrate concurrency and message passing between goroutines.
	// When: This line executes asynchronously; the main goroutine continues while this goroutine runs.
	// Where: Useful in real-world apps where tasks need to run concurrently without blocking the main program.

	msg := <-messages
	// What: Receives a string from the "messages" channel and assigns it to the variable "msg".
	// How: The "<-" operator reads data from the channel; if the channel is empty, it blocks until a value is available.
	// Why: To synchronize with the goroutine and retrieve the message it sent.
	// When: This line blocks execution until the goroutine writes to the channel.
	// Where: Channels often serve as a synchronization mechanism in concurrent programs.

	fmt.Println(msg)
	// What: Prints the value of "msg" to the standard output.
	// How: Uses the fmt.Println function to display the string followed by a newline.
	// Why: To show the result of the goroutine's computation or communication.
	// When: Happens after the message is successfully received from the channel.
	// Where: Used for debugging, logging, or displaying program results to the user.

}
