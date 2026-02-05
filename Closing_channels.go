package main

// ✅ What: Declares this file belongs to the main package.
// ✅ Why: Required for executable Go programs.
// ✅ When: Always at the top of a standalone Go file.
// ✅ Where: The Go runtime looks for main() here.

import "fmt"

// ✅ What: Imports the "fmt" package for formatted output.
// ✅ How: Provides fmt.Println to print messages.
// ✅ Why: Needed to display output.
// ✅ When: Required whenever printing or logging.

func main() {
	// ✅ What: Entry point of the program.
	// ✅ How: Executed automatically by Go runtime.
	// ✅ Why: Required in all executable Go programs.
	// ✅ When: Execution begins here.

	jobs := make(chan int, 5)
	// ✅ What: Creates a buffered channel of ints named jobs with capacity 5.
	// ✅ How: make(chan int, 5) allocates memory for buffered channel.
	// ✅ Why: Allows multiple jobs to be sent without immediate blocking.
	// ✅ When: Needed before sending jobs to a goroutine.
	// ✅ Where: Useful for producer-consumer or job queue patterns.

	done := make(chan bool)
	// ✅ What: Creates an unbuffered channel of bools named done.
	// ✅ How: make(chan bool) allocates memory for signaling completion.
	// ✅ Why: Used to notify main goroutine when all jobs are processed.
	// ✅ When: Needed before launching the worker goroutine.
	// ✅ Where: Commonly used for synchronization in Go concurrency.

	go func() {
		// ✅ What: Launches a goroutine to process jobs.
		// ✅ How: Anonymous function runs concurrently.
		// ✅ Why: Worker goroutine simulates job processing asynchronously.
		// ✅ When: Starts running while main continues to send jobs.

		for {
			j, more := <-jobs
			// ✅ What: Receives a job from jobs channel; also checks if channel is closed.
			// ✅ How: Two-value receive from a channel: j = value, more = bool indicating if channel is open.
			// ✅ Why: Detects when all jobs have been sent and channel is closed.
			// ✅ When: Loops indefinitely until channel is closed and drained.
			// ✅ Where: Pattern used in worker goroutines to process a stream of jobs safely.

			if more {
				fmt.Println("received job", j)
				// ✅ What: Prints the job if channel is still open.
				// ✅ How: Regular printing.
				// ✅ Why: Confirms job processing.
				// ✅ When: Each time a job is received.

			} else {
				fmt.Println("received all jobs")
				done <- true
				return
				// ✅ What: Channel is closed; signals main via done channel and exits goroutine.
				// ✅ How: send true to done channel; return terminates goroutine.
				// ✅ Why: Notifies main that processing is complete.
				// ✅ When: Happens once after all jobs are received.
				// ✅ Where: Proper cleanup and synchronization for goroutines.
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
		// ✅ What: Sends jobs 1, 2, 3 into jobs channel.
		// ✅ How: `<-` operator sends value into buffered channel.
		// ✅ Why: Populates the worker goroutine with tasks.
		// ✅ When: Happens sequentially in main goroutine.
		// ✅ Where: Producer stage in the producer-consumer pattern.
	}

	close(jobs)
	// ✅ What: Closes the jobs channel to signal no more jobs will be sent.
	// ✅ How: close(jobs) prevents further sends on the channel.
	// ✅ Why: Allows the worker to detect completion using the second value in receive (more).
	// ✅ When: After sending all jobs.
	// ✅ Where: Important to avoid deadlock or panics from sending on a closed channel.

	fmt.Println("sent all jobs")
	// ✅ What: Prints confirmation that all jobs were sent.
	// ✅ How: fmt.Println
	// ✅ Why: Logs progress.
	// ✅ When: After close(jobs).

	<-done
	// ✅ What: Blocks main goroutine until worker signals completion.
	// ✅ How: Receives value from done channel; discards value.
	// ✅ Why: Ensures main waits until all jobs are processed.
	// ✅ When: Synchronization point.

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
	// ✅ What: Tries to receive from closed jobs channel; ok will be false.
	// ✅ How: Two-value receive detects closed channel.
	// ✅ Why: Confirms that no more jobs are left in the channel.
	// ✅ When: Happens after done signal from worker.
	// ✅ Where: Demonstrates Go’s way to check if a channel is closed.

	// 🔹 Key insights:
	// 1. Buffered channels let main send multiple jobs without blocking.
	// 2. Closing a channel allows receivers to detect “end of data”.
	// 3. Two-value receive (`v, ok := <-ch`) is the idiomatic way to detect closed channels.
	// 4. done channel ensures main waits for goroutine completion, preventing premature exit.
}
