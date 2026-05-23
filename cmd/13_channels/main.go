package main

import "fmt"

func main(){

// In Go, a channel is a typed conduit through which goroutines communicate and synchronize. Instead of using explicit locks to access shared data, Go encourages "sharing memory by communicating" via these pipes. 


// Core Concepts
// Typed Conduit: Each channel is strictly typed to the data it carries (e.g., chan int for integers).
// The Arrow Operator (<-): Used for both sending and receiving:
// ch <- v: Send value v into channel ch.
// v := <-ch: Receive value from channel ch and assign it to v.
// Creation: You must initialize channels using the built-in make function; otherwise, they are nil and will block permanently.

	deadlockError()
	simpleExample()
	

}

func deadlockError(){

	c := make(chan int)
	// c <- 1
	// fmt.Println(c) //all goroutines are asleep - deadlock!
	
// This code causes a deadlock because unbuffered channels block until both a sender and a receiver are ready at the same time.

// The Mechanics of the Deadlock
// No Buffer: The line c := make(chan int) creates an unbuffered channel. It has zero capacity to store data.
// The Block: When the code reaches c <- 1, the main goroutine stops and waits for another goroutine to take the value 1.
// The Freeze: Because there are no other goroutines running to receive the data, the main goroutine will wait forever.
// The Panic: Go's runtime detects that all running goroutines are permanently asleep waiting for each other, so it crashes with a fatal error: all goroutines are asleep - deadlock!.

// The fix is 
	go func(){fmt.Println(<-c)}()
	c <- 1

// How above program works 
// Start: The main program launches a background goroutine.
// Wait: The goroutine reaches <-c and pauses because the channel is empty.
// Send: The main program continues and sends 1 into c.
// Transfer: The waiting goroutine instantly receives the 1.
// Print: The goroutine unblocks and prints 1.


}


func simpleExample(){
	messages := make(chan string) //creates a new channel that can carry string values.

    // Send a value into the channel from a new goroutine
    go func() { messages <- "ping" }()

    // Receive the value in the main goroutine
    msg := <-messages
    fmt.Println(msg)
}

// go func() { messages <- "ping" }()
// the () at the end of a function definition creates and immediately executes an Immediately Invoked Function Expression (IIFE).

// Normal way
// func printMessage() { ... }
// go printMessage() // Called using ()

// Anonymous Way
// go func() { ... }() // Defined and called all at once
