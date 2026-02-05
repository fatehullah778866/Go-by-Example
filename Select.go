package main

// ✅ What: Declares this file belongs to the main package.
// ✅ Why: Required for executable Go programs.
// ✅ When: Always at the top of a standalone Go program.
// ✅ Where: The main package allows the Go runtime to locate main().

import (
	"fmt"
	"time"
)

// ✅ What: Imports "fmt" for printing and "time" for time-related functions.
// ✅ How: "fmt" provides fmt.Println, "time" provides Sleep and duration types.
// ✅ Why: Needed for output and to simulate delays in goroutines.
// ✅ When: Required whenever using formatted output or timing functions.

func main() {
	// ✅ What: Entry point of the program.
	// ✅ How: Automatically executed by Go runtime.
	// ✅ Why: Required in every Go executable program.
	// ✅ When: Execution begins here.

	c1 := make(chan string)
	// ✅ What: Creates an unbuffered channel for strings named c1.
	// ✅ How: make(chan string) allocates the channel in memory.
	// ✅ Why: Used to receive messages from the first goroutine.
	// ✅ When: Needed before launching goroutines that send messages.
	// ✅ Where: Channels synchronize communication between goroutines.

	c2 := make(chan string)
	// ✅ What: Creates another unbuffered channel named c2.
	// ✅ How: Same as c1, used for the second goroutine.
	// ✅ Why: Separate channel allows independent communication streams.
	// ✅ When: Needed before starting goroutines that send messages.
	// ✅ Where: Useful in concurrent programming to handle multiple sources.

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	// ✅ What: Launches a goroutine that waits 1 second and sends "one" into c1.
	// ✅ How: Anonymous function runs concurrently; Sleep delays execution; send operator `<-` writes to channel.
	// ✅ Why: Simulates asynchronous work and delayed message sending.
	// ✅ When: Runs concurrently; main goroutine continues immediately.
	// ✅ Where: Models tasks that complete at different times in real-world apps.

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	// ✅ What: Launches a goroutine that waits 2 seconds and sends "two" into c2.
	// ✅ How: Same approach as the first goroutine but with a longer delay.
	// ✅ Why: Demonstrates how `select` handles whichever channel is ready first.
	// ✅ When: Runs concurrently; independent timing from c1.
	// ✅ Where: Shows multiple goroutines producing data at different speeds.

	for range 2 {
		// ✅ What: Loop runs 2 times, once for each expected message.
		// ✅ How: `range 2` is shorthand for iterating exactly twice.
		// ✅ Why: Ensures we handle all messages from c1 and c2.
		// ✅ When: Used after goroutines are launched to receive messages.
		// ✅ Where: Collects messages from multiple sources safely.

		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
			// ✅ What: Receives from c1 if ready and prints it.
			// ✅ How: select waits for a channel to have data; msg1 stores the value.
			// ✅ Why: Non-blocking way to handle multiple channels concurrently.
			// ✅ When: Executes immediately when c1 has a value.
			// ✅ Where: Typical in event-driven or concurrent pipelines.

		case msg2 := <-c2:
			fmt.Println("received", msg2)
			// ✅ What: Receives from c2 if ready and prints it.
			// ✅ How: Same as c1; select picks whichever channel is ready first.
			// ✅ Why: Ensures we don’t miss messages arriving at different times.
			// ✅ When: Executes immediately when c2 has a value.
			// ✅ Where: Useful in multiplexing multiple concurrent operations.
		}
	}

	// 🔹 Extra insight:
	// - `select` lets Go wait on **multiple channels simultaneously**.
	// - First message ready (from c1 after 1s) will be printed first, then the second (from c2 after 2s).
	// - This is a simple way to **handle asynchronous events** without busy-waiting.
	// - If no channel is ready, `select` **blocks** until one is.
}
