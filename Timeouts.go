package main

// ✅ What: Declares the main package for this program.
// ✅ Why: Required for an executable Go program.
// ✅ When: Always at the top of a standalone Go file.
// ✅ Where: The main package allows Go runtime to locate main().

import (
	"fmt"
	"time"
)

// ✅ What: Imports "fmt" for printing and "time" for timing functions.
// ✅ How: "fmt" provides fmt.Println; "time" provides Sleep and After.
// ✅ Why: Needed to simulate delays and implement timeouts.
// ✅ When: Required whenever output or time-based operations are used.

func main() {
	// ✅ What: Entry point of the program.
	// ✅ How: Go runtime automatically executes main().
	// ✅ Why: Required in all executable Go programs.
	// ✅ When: Execution starts here.

	c1 := make(chan string, 1)
	// ✅ What: Creates a buffered channel of strings with capacity 1.
	// ✅ How: make(chan string, 1) allocates memory for a buffered channel.
	// ✅ Why: Buffered channel allows one value to be sent without blocking.
	// ✅ When: Needed before sending messages from goroutine.

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()
	// ✅ What: Launches a goroutine that waits 2 seconds and sends "result 1" into c1.
	// ✅ How: Anonymous function runs concurrently; time.Sleep simulates work.
	// ✅ Why: Demonstrates asynchronous operation that may take longer than a timeout.
	// ✅ When: Goroutine runs independently of main function.
	// ✅ Where: Useful in simulating long-running tasks in concurrent programs.

	select {
	case res := <-c1:
		fmt.Println(res)
		// ✅ What: Receives from c1 if ready and prints the result.
		// ✅ How: `<-c1` reads from channel; stored in res.
		// ✅ Why: Handles message if task finishes before timeout.
		// ✅ When: Executes immediately if c1 has a value.
		// ✅ Where: Useful in pipelines or async message handling.

	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
		// ✅ What: Receives a value from time.After channel after 1 second.
		// ✅ How: time.After returns a channel that sends current time after duration.
		// ✅ Why: Implements a timeout mechanism; prevents blocking forever.
		// ✅ When: Executes if c1 does not provide a value within 1 second.
		// ✅ Where: Typical in real-world programs with network requests or long tasks.
	}

	c2 := make(chan string, 1)
	// ✅ What: Creates another buffered channel of strings with capacity 1.
	// ✅ How: Same as c1, allocates memory for channel.
	// ✅ Why: Separate channel for the second task.
	// ✅ When: Needed before launching the next goroutine.

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	// ✅ What: Launches a goroutine that waits 2 seconds and sends "result 2" into c2.
	// ✅ How: Simulates a delayed computation asynchronously.
	// ✅ Why: Demonstrates a task that may complete **before timeout** this time.
	// ✅ When: Goroutine runs concurrently.

	select {
	case res := <-c2:
		fmt.Println(res)
		// ✅ What: Receives the value from c2 and prints it.
		// ✅ How: `<-c2` reads from channel; fmt.Println prints it.
		// ✅ Why: Handles the result if the goroutine finishes before timeout.
		// ✅ When: Executes immediately if c2 has a value ready.
		// ✅ Where: Pattern for waiting on a result with optional timeout.

	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
		// ✅ What: Receives a value from time.After channel after 3 seconds.
		// ✅ How: time.After returns a channel that fires once after the duration.
		// ✅ Why: Implements timeout mechanism to avoid blocking forever.
		// ✅ When: Executes if c2 does not provide a value within 3 seconds.
		// ✅ Where: Ensures program does not hang on slow operations.
	}

	// 🔹 Key insights:
	// 1. `time.After(duration)` creates a channel that sends a timestamp after the duration.
	// 2. Using `select`, Go can **wait on multiple channels simultaneously**, including timeout channels.
	// 3. In this example:
	//    - First select triggers the timeout because c1 sleeps 2s but timeout is 1s → prints "timeout 1".
	//    - Second select completes successfully because c2 sleeps 2s but timeout is 3s → prints "result 2".
	// 4. This pattern is common for **network calls, I/O, or any operation where you want a fallback if it takes too long**.
}
