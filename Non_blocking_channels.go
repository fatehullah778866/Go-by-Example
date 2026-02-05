package main

// ✅ What: Declares the main package for this program.
// ✅ Why: Required for any executable Go program.
// ✅ When: Always at the top of a standalone Go file.
// ✅ Where: The Go runtime looks here for main() function.

import "fmt"

// ✅ What: Imports "fmt" package for formatted input/output.
// ✅ How: Provides fmt.Println to print messages to the console.
// ✅ Why: Needed to display output.
// ✅ When: Required whenever printing or logging.

func main() {
	// ✅ What: Entry point of the program.
	// ✅ How: Automatically executed by the Go runtime.
	// ✅ Why: Required for all executable Go programs.
	// ✅ When: Execution begins here.

	messages := make(chan string)
	// ✅ What: Creates an unbuffered channel for strings named messages.
	// ✅ How: make(chan string) allocates memory for the channel.
	// ✅ Why: Used for sending and receiving messages between goroutines.
	// ✅ When: Needed before any send/receive operations on this channel.
	// ✅ Where: Channels are Go’s primary concurrency primitive.

	signals := make(chan bool)
	// ✅ What: Creates an unbuffered channel for bool values named signals.
	// ✅ How: Same as messages, allocates memory for channel.
	// ✅ Why: Used to simulate signals or notifications between goroutines.
	// ✅ When: Needed before send/receive operations on this channel.
	// ✅ Where: Can be used to coordinate state or events.

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
		// ✅ What: Tries to receive a message from messages channel.
		// ✅ How: `<-messages` reads from channel; msg stores the value.
		// ✅ Why: Demonstrates non-blocking receive with select.
		// ✅ When: Executes if a message is ready in the channel.
		// ✅ Where: Useful to check for a message without blocking.
	default:
		fmt.Println("no message received")
		// ✅ What: Executes if no channel in select is ready.
		// ✅ How: default ensures select doesn’t block.
		// ✅ Why: Prevents program from waiting indefinitely for a message.
		// ✅ When: Executes immediately if messages channel is empty.
		// ✅ Where: Pattern for non-blocking reads.

	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
		// ✅ What: Tries to send msg into messages channel.
		// ✅ How: `<-` operator sends the value into channel.
		// ✅ Why: Demonstrates non-blocking send with select.
		// ✅ When: Executes only if channel can accept value immediately.
		// ✅ Where: Useful to attempt sending without blocking the main goroutine.
	default:
		fmt.Println("no message sent")
		// ✅ What: Executes if messages channel is not ready for send.
		// ✅ How: default prevents blocking.
		// ✅ Why: Ensures program continues instead of waiting for a receiver.
		// ✅ When: Happens when channel is unbuffered and no receiver is waiting.
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
		// ✅ What: Tries to receive from messages channel.
		// ✅ How: `<-messages` reads value if available.
		// ✅ Why: Handles the message if any was previously sent.
		// ✅ When: Executes only if channel has a message.

	case sig := <-signals:
		fmt.Println("received signal", sig)
		// ✅ What: Tries to receive a signal from signals channel.
		// ✅ How: `<-signals` reads value if available.
		// ✅ Why: Demonstrates multiple channels in a select.
		// ✅ When: Executes only if channel has a signal ready.

	default:
		fmt.Println("no activity")
		// ✅ What: Executes if none of the channels are ready.
		// ✅ How: default prevents blocking.
		// ✅ Why: Ensures program continues immediately even if no messages/signals.
		// ✅ When: Happens when both channels are empty/unready.
		// ✅ Where: Useful for polling channels or doing non-blocking checks.
	}

	// 🔹 Key insights:
	// 1. Using `default` in a select makes it **non-blocking**.
	// 2. First select prints "no message received" because messages is empty.
	// 3. Second select prints "no message sent" because no receiver is waiting on unbuffered channel.
	// 4. Third select prints "no activity" because both messages and signals channels are empty.
	// 5. This pattern is useful for **polling, timeouts, or optional message handling**.
}
