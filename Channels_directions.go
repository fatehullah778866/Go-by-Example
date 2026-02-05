package main

// ✅ What: Declares the main package for this program.
// ✅ Why: Required for an executable Go program.
// ✅ When: At the top of the Go file.
// ✅ Where: Must be "main" so the Go runtime can find main().

import "fmt"

// ✅ What: Imports the "fmt" package for formatted output.
// ✅ How: Provides functions like fmt.Println to print to console.
// ✅ Why: Needed to display the final message.
// ✅ When: Whenever output is needed to console.

func ping(pings chan<- string, msg string) {
	// ✅ What: Defines a function `ping` that sends a string into a channel.
	// ✅ How: The parameter `chan<- string` is a **send-only channel**.
	// ✅ Why: Restricts this function to only send messages, not receive.
	// ✅ When: Called when we want to send a message to another goroutine or process.
	// ✅ Where: Useful in concurrency patterns to prevent accidental reads in a producer function.

	pings <- msg
	// ✅ What: Sends `msg` into the `pings` channel.
	// ✅ How: Uses `<-` operator to write into the channel.
	// ✅ Why: Allows the message to be transferred to the consumer (pong function).
	// ✅ When: Blocks only if the channel buffer is full (here buffer=1, so non-blocking).
	// ✅ Where: This is the “sending” part of the ping-pong mechanism.
}

func pong(pings <-chan string, pongs chan<- string) {
	// ✅ What: Defines `pong` function that receives from one channel and sends to another.
	// ✅ How: `pings <-chan string` is a **receive-only channel**, `pongs chan<- string` is **send-only**.
	// ✅ Why: Clear separation of responsibilities: `pong` only reads from pings and writes to pongs.
	// ✅ When: Called after ping has sent the message.
	// ✅ Where: Common in pipelines where stages read from one channel and write to another.

	msg := <-pings
	// ✅ What: Receives a message from the `pings` channel.
	// ✅ How: `<-pings` reads from the channel; blocks if the channel is empty.
	// ✅ Why: Synchronizes with the ping function; ensures message has arrived before proceeding.
	// ✅ When: Blocks until a message is available.
	// ✅ Where: Reading from channels enforces safe communication between concurrent processes.

	pongs <- msg
	// ✅ What: Sends the received message into the `pongs` channel.
	// ✅ How: `<-` operator writes into the send-only channel.
	// ✅ Why: Passes the message along to the next stage or consumer.
	// ✅ When: Blocks if pongs channel buffer is full; here buffer=1, so non-blocking.
	// ✅ Where: Enables chaining or pipelining of messages between goroutines.
}

func main() {
	// ✅ What: Entry point of the program.
	// ✅ How: Go runtime calls main() automatically.
	// ✅ Why: Required in executable Go programs.
	// ✅ When: First function executed when program starts.

	pings := make(chan string, 1)
	// ✅ What: Creates a buffered channel for sending messages to ping.
	// ✅ How: Buffer size 1 allows one message to be sent without blocking.
	// ✅ Why: Ensures ping can send message even if pong hasn't read yet.
	// ✅ When: Used for message passing between functions.
	// ✅ Where: Producer-consumer or pipeline pattern.

	pongs := make(chan string, 1)
	// ✅ What: Creates a buffered channel for receiving messages from pong.
	// ✅ How: Buffer size 1 allows pong to send message without blocking immediately.
	// ✅ Why: Needed to pass the message from pong to main function.
	// ✅ When: Part of the ping-pong communication setup.
	// ✅ Where: Acts as the final channel for message consumption.

	ping(pings, "passed message")
	// ✅ What: Calls ping function to send a message.
	// ✅ How: Passes the pings channel and the string "passed message".
	// ✅ Why: Initiates the sending stage of the ping-pong pipeline.
	// ✅ When: Happens before pong reads from the channel.
	// ✅ Where: First step in ping-pong message flow.

	pong(pings, pongs)
	// ✅ What: Calls pong function to receive from pings and send to pongs.
	// ✅ How: Passes channels for input (pings) and output (pongs).
	// ✅ Why: Demonstrates channel communication between functions.
	// ✅ When: After ping has sent a message, this function receives and forwards it.
	// ✅ Where: Middle stage of the ping-pong pipeline.

	fmt.Println(<-pongs)
	// ✅ What: Receives the message from pongs channel and prints it.
	// ✅ How: `<-pongs` reads from the channel; fmt.Println prints it.
	// ✅ Why: Shows the final result of the ping-pong communication.
	// ✅ When: Blocks until pong sends the message (synchronization point).
	// ✅ Where: Consumer stage in the ping-pong pipeline.

	// 🔹 Extra insight:
	// The use of **directional channels** (`chan<-` and `<-chan`) ensures **type safety**:
	//   - ping can only send.
	//   - pong reads from pings and sends to pongs.
	// This prevents accidental misuse and makes pipeline logic explicit.
}
