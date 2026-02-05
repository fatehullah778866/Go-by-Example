package main

// ✅ What: Declares the main package for this program.
// ✅ Why: Required for executable Go programs.
// ✅ When: Always at the top of a standalone Go file.
// ✅ Where: The Go runtime looks for main() here.

import "fmt"

// ✅ What: Imports "fmt" package for formatted input/output.
// ✅ How: Provides fmt.Println to print messages to console.
// ✅ Why: Needed to display output.
// ✅ When: Required whenever printing or logging.

func main() {
	// ✅ What: Entry point of the program.
	// ✅ How: Automatically executed by Go runtime.
	// ✅ Why: Required for all executable Go programs.
	// ✅ When: Execution begins here.

	queue := make(chan string, 2)
	// ✅ What: Creates a buffered channel of strings named queue with capacity 2.
	// ✅ How: make(chan string, 2) allocates memory for the buffered channel.
	// ✅ Why: Buffered channel allows sending multiple values without immediate blocking.
	// ✅ When: Needed before sending elements into the channel.
	// ✅ Where: Acts as a small queue for storing messages.

	queue <- "one"
	// ✅ What: Sends the string "one" into the queue channel.
	// ✅ How: `<-` operator writes the value into the buffered channel.
	// ✅ Why: Adds the first element to the queue.
	// ✅ When: Non-blocking because the buffer has capacity.
	// ✅ Where: Producer adds data into the channel.

	queue <- "two"
	// ✅ What: Sends the string "two" into the queue channel.
	// ✅ How: Same `<-` operator fills the buffer.
	// ✅ Why: Adds the second element to the queue.
	// ✅ When: Non-blocking because buffer size = 2.
	// ✅ Where: Producer stage continues.

	close(queue)
	// ✅ What: Closes the queue channel to signal no more values will be sent.
	// ✅ How: close(queue) prevents future sends.
	// ✅ Why: Enables safe iteration over the channel with `range`.
	// ✅ When: After all values are sent.
	// ✅ Where: Closing a channel is essential for `for elem := range queue` loops.

	for elem := range queue {
		fmt.Println(elem)
		// ✅ What: Iterates over each element in the queue until the channel is closed.
		// ✅ How: `range` automatically receives values until the channel is drained and closed.
		// ✅ Why: Provides a clean, idiomatic way to process all values in a channel.
		// ✅ When: Loop executes once per element; terminates automatically after channel is empty and closed.
		// ✅ Where: Consumer stage in producer-consumer pattern.
	}

	// 🔹 Key insights:
	// 1. Buffered channels allow multiple elements to be queued without blocking immediately.
	// 2. Closing the channel signals the end of data; `range` can safely iterate until done.
	// 3. Using `for elem := range channel` is an idiomatic Go pattern for consuming all channel messages.
}
