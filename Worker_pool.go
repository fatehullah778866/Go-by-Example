package main

// ✅ What: Declares the main package for this program.
// ✅ Why: Required for executable Go programs.
// ✅ When: At the top of any standalone Go file.
// ✅ Where: The Go runtime looks for main() here.

import (
	"fmt"
	"time"
)

// ✅ What: Imports "fmt" for printing and "time" for simulating work duration.
// ✅ How: fmt.Println prints messages; time.Sleep pauses execution for a duration.
// ✅ Why: Needed to display output and simulate processing delays.
// ✅ When: Required whenever printing or delaying execution.

func worker(id int, jobs <-chan int, results chan<- int) {
	// ✅ What: Worker function that processes jobs and sends results.
	// ✅ How: Receives jobs from jobs channel, performs computation, sends results.
	// ✅ Why: Implements concurrent workers in a worker pool pattern.
	// ✅ When: Called as a goroutine for each worker.
	// ✅ Where: Typical pattern in Go for parallel task processing.

	for j := range jobs {
		// ✅ What: Iterates over jobs channel until closed.
		// ✅ How: `range` automatically stops when channel is closed.
		// ✅ Why: Worker continuously processes jobs as long as they are available.
		// ✅ When: Runs as long as main sends jobs.
		// ✅ Where: Consumer in producer-consumer pattern.

		fmt.Println("worker", id, "started  job", j)
		// ✅ What: Prints start of job processing.
		// ✅ How: fmt.Println with worker id and job number.
		// ✅ Why: Logs progress for clarity.
		// ✅ When: Before processing each job.

		time.Sleep(time.Second)
		// ✅ What: Simulates time-consuming work (1 second per job).
		// ✅ How: time.Sleep pauses execution of the current goroutine.
		// ✅ Why: Demonstrates that workers can process jobs concurrently.
		// ✅ When: Each job takes some simulated time.

		fmt.Println("worker", id, "finished job", j)
		// ✅ What: Prints completion of job processing.
		// ✅ How: fmt.Println with worker id and job number.
		// ✅ Why: Logs that job processing is done.

		results <- j * 2
		// ✅ What: Sends the result of the job (double the number) to results channel.
		// ✅ How: `<-results` writes the computed value.
		// ✅ Why: Worker communicates output back to main.
		// ✅ When: After job is processed.
		// ✅ Where: Producer side of results channel.
	}
}

func main() {
	// ✅ What: Entry point of the program.
	// ✅ How: Automatically executed by Go runtime.
	// ✅ Why: Required in all executable Go programs.

	const numJobs = 5
	// ✅ What: Defines total number of jobs to process.
	// ✅ Why: Controls size of jobs channel and iterations.
	// ✅ When: Used to initialize channels and loops.

	jobs := make(chan int, numJobs)
	// ✅ What: Buffered channel for sending jobs to workers.
	// ✅ How: Capacity = numJobs allows all jobs to be sent without blocking.
	// ✅ Why: Buffered channels improve efficiency in worker pools.
	// ✅ When: Before sending jobs.

	results := make(chan int, numJobs)
	// ✅ What: Buffered channel for receiving processed results from workers.
	// ✅ Why: Stores results for main to consume later.
	// ✅ When: Before starting workers.

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
		// ✅ What: Launches 3 worker goroutines.
		// ✅ How: Each worker processes jobs from the jobs channel concurrently.
		// ✅ Why: Allows parallel processing for efficiency.
		// ✅ When: Before sending jobs.
		// ✅ Where: Typical worker pool setup.
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
		// ✅ What: Sends jobs 1 to 5 into jobs channel.
		// ✅ How: `<-jobs` writes each number into buffered channel.
		// ✅ Why: Populates worker pool with tasks.
		// ✅ When: After workers are started.
		// ✅ Where: Producer side of jobs channel.
	}

	close(jobs)
	// ✅ What: Closes the jobs channel to signal no more jobs will be sent.
	// ✅ Why: Workers stop processing after draining the channel.
	// ✅ When: After all jobs are sent.
	// ✅ Where: Proper cleanup in worker pool pattern.

	for a := 1; a <= numJobs; a++ {
		<-results
		// ✅ What: Receives results from results channel.
		// ✅ How: `<-results` blocks until a value is available.
		// ✅ Why: Ensures main waits until all jobs are processed.
		// ✅ When: Iterates exactly numJobs times.
		// ✅ Where: Consumer side of results channel.
	}

	// 🔹 Key insights:
	// 1. Worker pool allows concurrent job processing using goroutines.
	// 2. Buffered channels for jobs/results prevent blocking during sending.
	// 3. Closing jobs channel signals workers that no more jobs are coming.
	// 4. Receiving from results ensures main waits for all job completions.
	// 5. This pattern scales easily by adding more workers or jobs.
}
