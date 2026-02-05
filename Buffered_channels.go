package main

// What: Declares this file belongs to the main package.
// Why: "main" package is required for an executable Go program.
// When: Every standalone Go program needs a main package.
// Where: At the very top of the Go file.

import "fmt"

// What: Imports the "fmt" package for formatted input/output.
// How: Provides functions like fmt.Println to print to console.
// Why: Necessary to output the content of the channel.
// When: Needed whenever we want to display output.

func main() {
	// What: Defines the entry point of the program.
	// How: Program execution starts from this function.
	// Why: Required in every Go executable.
	// When: Called automatically when the program starts.

	messages := make(chan string, 2)
	// What: Creates a buffered channel of type string with capacity 2.
	// How: The make() function allocates a channel that can hold up to 2 elements.
	// Why: Buffered channels allow sending multiple messages without immediately blocking the sender.
	// When: Use buffered channels when you want to decouple sender and receiver timing.
	// Where: Buffered channels are useful in producer-consumer patterns or concurrent pipelines.

	messages <- "buffered"
	// What: Sends the string "buffered" into the messages channel.
	// How: The "<-" operator writes the value into the channel.
	// Why: To demonstrate sending multiple messages into a buffered channel.
	// When: This does NOT block because the channel has space (capacity 2).
	// Where: Sending into a buffered channel allows temporary storage of data before receiving.

	messages <- "channel"
	// What: Sends the string "channel" into the messages channel.
	// How: Same "<-" operator writes the value into the channel.
	// Why: Fills the second slot of the buffered channel.
	// When: Still non-blocking because the channel's buffer is not full.
	// Where: Buffered channels enable multiple messages to queue up for later retrieval.

	fmt.Println(<-messages)
	// What: Receives the first string from the messages channel and prints it.
	// How: "<-messages" reads from the channel; fmt.Println outputs it.
	// Why: To retrieve messages in the order they were sent (FIFO).
	// When: This blocks if the channel were empty, but here the buffer has data.
	// Where: Buffered channels maintain order, making them predictable for consumers.

	fmt.Println(<-messages)
	// What: Receives the second string from the messages channel and prints it.
	// How: Same reading and printing process.
	// Why: Demonstrates that all buffered messages are accessible in the correct order.
	// When: Blocks if the buffer were empty; here it prints the remaining message.
	// Where: Useful in scenarios where multiple producers send data and a consumer reads them sequentially.

}
