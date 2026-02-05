package main

// ✅ What: Declares the main package for this program.
// ✅ Why: Required for any executable Go program.
// ✅ When: At the top of every standalone Go file.
// ✅ Where: The Go runtime looks here for main().

import (
	"fmt"
	"time"
)

// ✅ What: Imports "fmt" for printing and "time" for timer functions.
// ✅ How: fmt.Println prints messages; time provides NewTimer, Sleep, and time.Duration.
// ✅ Why: Needed to create and control timers, and to display output.
// ✅ When: Required whenever working with timing or delayed execution.

func main() {
	// ✅ What: Entry point of the program.
	// ✅ How: Automatically executed by the Go runtime.
	// ✅ Why: Required for all executable Go programs.
	// ✅ When: Execution begins here.

	timer1 := time.NewTimer(2 * time.Second)
	// ✅ What: Creates a timer that will send the current time on its channel after 2 seconds.
	// ✅ How: time.NewTimer returns a *Timer, which has a channel C that fires once.
	// ✅ Why: Demonstrates a one-shot timer in Go.
	// ✅ When: Starts counting down immediately upon creation.
	// ✅ Where: Useful for delayed execution or timeouts.

	<-timer1.C
	// ✅ What: Receives from the timer's channel, effectively blocking until the timer fires.
	// ✅ How: `<-timer1.C` waits for the timer to send the current time.
	// ✅ Why: Blocks main execution until the timer expires.
	// ✅ When: Happens 2 seconds after timer1 creation.
	// ✅ Where: Synchronizes execution after the delay.

	fmt.Println("Timer 1 fired")
	// ✅ What: Prints that timer1 has fired.
	// ✅ How: Simple fmt.Println call.
	// ✅ Why: Confirms timer expiration.
	// ✅ When: Immediately after receiving from timer1.C.
	// ✅ Where: Logs events in the program.

	timer2 := time.NewTimer(time.Second)
	// ✅ What: Creates a second timer that fires after 1 second.
	// ✅ How: time.NewTimer returns a *Timer with channel C.
	// ✅ Why: Demonstrates stopping a timer before it fires.
	// ✅ When: Starts counting down immediately.

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// ✅ What: Launches a goroutine that waits for timer2 to fire and prints a message.
	// ✅ How: Anonymous goroutine blocks on `<-timer2.C`.
	// ✅ Why: Simulates asynchronous execution with a timer.
	// ✅ When: Could execute after 1 second if timer is not stopped.
	// ✅ Where: Shows how timers work with goroutines.

	stop2 := timer2.Stop()
	// ✅ What: Stops timer2 before it fires.
	// ✅ How: timer2.Stop() returns true if the timer was active and successfully stopped.
	// ✅ Why: Prevents timer2 from firing and executing the goroutine.
	// ✅ When: Called immediately after goroutine launch.
	// ✅ Where: Useful to cancel delayed tasks in real applications.

	if stop2 {
		fmt.Println("Timer 2 stopped")
		// ✅ What: Prints confirmation that timer2 was stopped successfully.
		// ✅ How: Conditional on timer2.Stop() return value.
		// ✅ Why: Provides feedback that the timer was canceled before firing.
		// ✅ When: Executed immediately if timer was stopped before firing.
		// ✅ Where: Logging or controlling program flow.
	}

	time.Sleep(2 * time.Second)
	// ✅ What: Sleeps main goroutine for 2 seconds to allow any pending timers/goroutines to complete.
	// ✅ How: time.Sleep blocks the main goroutine for the specified duration.
	// ✅ Why: Ensures program does not exit immediately, giving time to observe behavior.
	// ✅ When: After all timer logic.
	// ✅ Where: Common in examples to prevent premature program exit.

	// 🔹 Key insights:
	// 1. time.NewTimer creates a one-shot timer that fires once after a duration.
	// 2. Receiving from timer.C blocks until the timer fires.
	// 3. timer.Stop() can cancel a timer if it hasn’t fired yet.
	// 4. Goroutines combined with timers allow asynchronous, delayed execution.
	// 5. time.Sleep at the end ensures main doesn’t exit before observing timers.
}
