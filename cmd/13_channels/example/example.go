package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5

func main() {
	chickenChannel := make(chan (string))
	websites := []string{"walmart.com", "costco.com", "target.com", "amazon.com"}
	for _, website := range websites {
		go checkChickenPrice(chickenChannel, website)
	}
	sendMessage(chickenChannel)

}

func checkChickenPrice(c chan (string), website string) {
	for {
		fmt.Printf("%v entered the chat\n", website)
		time.Sleep(time.Second * 1)
		chickenPrice := rand.Float32() * 20
		fmt.Printf("Chicken price is %v in %v\n", chickenPrice, website)
		if chickenPrice <= MAX_CHICKEN_PRICE {
			fmt.Println("Got the price")
			c <- website
			break
		}

	}
}

func sendMessage(c chan (string)) {
	fmt.Printf("The message has been send to the %v website.\n", <-c)
}


// This program uses Go goroutines and an unbuffered channel to concurrently search for the cheapest chicken price across four websites.
// Concurrent Spawning: The main function starts four separate background workers (one for each website) to check prices at the same time.
// Asynchronous Looping: Each website worker loops independently, sleeps for 1 second, and generates a random float price between 0.0 and 20.0.
// Main Thread Blocking: The main function calls sendMessage, which stops and waits indefinitely for the first value to arrive on the channel.
// The Winner: The moment a website finds a price less than or equal to $5 (costco.com at 2.57), it breaks its loop and sends its name into the channel.
// Instant Shutdown: The main function receives this first name, prints the final message, and immediately terminates the program—instantly killing any other websites still sleeping or processing in the background.