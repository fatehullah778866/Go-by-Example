package main

// ✅ What: Declares the main package for this program.
// ✅ Why: Required for executable Go programs.
// ✅ When: Always at the top of a standalone Go file.
// ✅ Where: The Go runtime looks for main() here.

import (
	"fmt"
	"time"
)

// ✅ What: Imports "fmt" for printing and "time" for ticker and timing functions.
// ✅ How: fmt.Println prints output; time provides NewTicker, Sleep, and durations.
// ✅ Why: Needed to create periodic events and display output.
// ✅ When: Required whenever printing or working with time-based events.

func main() {
	// ✅ What: Entry point of the program.
	// ✅ How: Automatically executed by Go runtime.
	// ✅ Why: Required in all executable Go programs.
	// ✅ When: Execution begins here.

	ticker := time.NewTicker(500 * time.Millisecond)
	// ✅ What: Creates a ticker that sends the current time every 500 milliseconds.
	// ✅ How: time.NewTicker returns a *Ticker with a channel C that ticks repeatedly.
	// ✅ Why: Demonstrates periodic events in Go.
	// ✅ When: Starts ticking immediately upon creation.
	// ✅ Where: Useful for repeated actions at regular intervals.

	done := make(chan bool)
	// ✅ What: Creates an unbuffered channel to signal when to stop the ticker goroutine.
	// ✅ How: make(chan bool) allocates memory for signaling.
	// ✅ Why: Allows main to tell the goroutine to exit safely.
	// ✅ When: Needed before launching the goroutine.
	// ✅ Where: Synchronization channel for safe shutdown.

	go func() {
		for {
			select {
			case <-done:
				return
				// ✅ What: Exits the goroutine when done channel receives a value.
				// ✅ How: `<-done` blocks until main sends a signal; return exits the loop.
				// ✅ Why: Gracefully stops the ticker goroutine.
				// ✅ When: Triggered after main decides to stop the ticker.
				// ✅ Where: Proper way to stop a long-running goroutine.

			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				// ✅ What: Receives a tick from ticker.C and prints the current time.
				// ✅ How: `<-ticker.C` blocks until the next tick; stored in t.
				// ✅ Why: Demonstrates periodic execution of code.
				// ✅ When: Fires every 500 milliseconds until stopped.
				// ✅ Where: Common in scheduling tasks or heartbeats.
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	// ✅ What: Sleeps main goroutine for 1.6 seconds to allow ticker goroutine to run.
	// ✅ How: time.Sleep blocks execution for the specified duration.
	// ✅ Why: Ensures the program doesn’t exit immediately; lets ticks happen.
	// ✅ When: After starting the ticker goroutine.
	// ✅ Where: Useful in examples to simulate program runtime.

	ticker.Stop()
	// ✅ What: Stops the ticker from sending further ticks.
	// ✅ How: ticker.Stop() prevents more sends on ticker.C.
	// ✅ Why: Proper cleanup of resources and stopping periodic events.
	// ✅ When: After 1.6 seconds, to prevent further ticks.
	// ✅ Where: Best practice to stop tickers when no longer needed.

	done <- true
	// ✅ What: Sends a signal to the goroutine to exit the loop.
	// ✅ How: `<-done` writes true into the done channel.
	// ✅ Why: Ensures the goroutine terminates gracefully.
	// ✅ When: After stopping the ticker.
	// ✅ Where: Coordination between main and goroutine for safe shutdown.

	fmt.Println("Ticker stopped")
	// ✅ What: Prints confirmation that ticker has been stopped.
	// ✅ How: fmt.Println prints output.
	// ✅ Why: Logs that the periodic execution has ended.
	// ✅ When: After sending stop signal and stopping ticker.
	// ✅ Where: Useful for tracking program flow.

	// 🔹 Key insights:
	// 1. time.NewTicker creates a **repeated timer** that ticks at fixed intervals.
	// 2. select in goroutine handles multiple channels: ticker.C for ticks, done for shutdown.
	// 3. ticker.Stop() halts further ticks and prevents resource leaks.
	// 4. Using done channel ensures **graceful goroutine termination**.
	// 5. This pattern is useful for periodic tasks, heartbeats, or polling events in Go.
}
