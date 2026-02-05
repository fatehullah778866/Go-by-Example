package main

// ✅ What: Declares the main package for the program.
// ✅ Why: Needed for any executable Go program; defines entry point.
// ✅ When: At the top of every Go file meant to be run as a program.
// ✅ Where: Must be "main" to allow the Go runtime to find main().

import "fmt"

// ✅ What: Imports the "fmt" package for formatted output.
// ✅ How: Provides functions like fmt.Println for printing strings to console.
// ✅ Why: Without fmt, we cannot display messages or debug output.
// ✅ When: Always required when printing/logging in console apps.

func main() {
	// ✅ What: The main function is the entry point for program execution.
	// ✅ How: Go runtime calls main() automatically when program starts.
	// ✅ Why: Required in the "main" package to run any Go program.
	// ✅ When: Called first when program executes.

	messages := make(chan string, 2)
	// ✅ What: Creates a buffered channel of strings with capacity 2.
	// ✅ How: make() allocates memory and initializes the channel with a fixed buffer size.
	// ✅ Why: Buffered channels allow sending multiple values without blocking immediately.
	// ✅ When: Use when you want to decouple sender and receiver timing.
	// ✅ Where: Ideal for producer-consumer patterns or pipelined goroutines.

	messages <- "buffered"
	// ✅ What: Sends "buffered" into the channel.
	// ✅ How: "<-" operator writes into the channel buffer.
	// ✅ Why: Demonstrates sending a message into a buffered channel.
	// ✅ When: Does not block because the buffer has free space (1/2 used now).
	// ✅ Where: Buffered channels allow temporary storage before consumption.

	messages <- "channel"
	// ✅ What: Sends "channel" into the channel.
	// ✅ How: Same "<-" operator fills the second slot of the buffer.
	// ✅ Why: Shows multiple messages can be queued in a buffered channel.
	// ✅ When: Still non-blocking because buffer capacity is 2 (now full).
	// ✅ Where: Buffered channels queue messages for later retrieval.

	fmt.Println(<-messages)
	// ✅ What: Receives the first string from the channel and prints it.
	// ✅ How: "<-messages" reads the value; fmt.Println outputs it.
	// ✅ Why: Demonstrates FIFO behavior: first sent, first received.
	// ✅ When: Blocks only if channel were empty; here, buffer has data.
	// ✅ Where: Buffered channels maintain predictable order for consumers.

	fmt.Println(<-messages)
	// ✅ What: Receives the second string from the channel and prints it.
	// ✅ How: Same read and print process.
	// ✅ Why: Shows all buffered messages can be accessed sequentially.
	// ✅ When: Blocks if channel is empty; here, prints remaining message.
	// ✅ Where: Buffered channels support multiple queued messages safely.

	// 🔹 Extra insight:
	// If we tried to send a third message here (messages <- "extra"),
	// it would **block** until a message was received from the channel, because the buffer is full.
	// This is how Go enforces synchronization even with buffered channels.
}
