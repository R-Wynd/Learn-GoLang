package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

var wg = sync.WaitGroup{} //w1. Define the WaitGroup
var mu sync.Mutex // m1. Define the Mutex

func main(){

	// Concept #1: Go routines
	// A goroutine is a lightweight thread managed by the Go runtime. It allows your program to run multiple functions concurrently (at the same time) with incredibly low memory overhead


	// Concept #2: WaitGroups
	// Using time.Sleep() to wait for goroutines is bad practice because you never know exactly how long a task will take. Instead, Go provides sync.WaitGroup to coordinate them. 

	// Concept #3: Mutex
	// A Mutex (short for Mutual Exclusion) is a synchronization tool used to prevent race conditions. It ensures that only one goroutine can access a specific block of code or a variable at any given time.
	// Think of a Mutex as a digital bathroom key in a coffee shop. If someone has the key, they are inside using the bathroom. If you want to use it, you must wait in line until they come out and hand over the key.

	// Concept #4: ReadMutex
	// A ReadLock (provided by sync.RWMutex in Go) is an optimized type of lock used when you have a piece of data that gets read frequently by many goroutines, but modified infrequently.

	
	t0 := time.Now()
	for i := 0; i<len(dbData); i++{
		
		wg.Add(1)  //w2. Tell the WaitGroup to expect 1 more worker
		go dbCall(i)
		
	}
	wg.Wait() // w4. Tell the waitGroup to wait till it comes again to 0 
	fmt.Printf("The results are %v\n", result)
	fmt.Printf("The total execution time is: %v\n", time.Since(t0))

}

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var result = []string{}

func dbCall(i int){
	fmt.Println("Called this", i)
	delay := rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Printf("The result from the database is: %v\n", dbData[i])


	mu.Lock()  //m2. Locks -> Grab the key. Other goroutines must wait here.
	result = append(result, dbData[i])  //m3. One go routine at a time can access (Safe to modify)
	mu.Unlock()   // 4. Return the key.

	wg.Done() //w3. Tell the WaitGroup this worker is done and it decrements the waitGroup
}


// Output 
// Called this 4
// Called this 2
// Called this 0
// Called this 1
// Called this 3
// ^^It call all the goroutines <- go


// The result from the database is: id5
// The result from the database is: id1
// The result from the database is: id3
// The result from the database is: id2
// The result from the database is: id4
//^^Each go routines Enters into func and decrement 1 value <- wg.Done)

// The total execution time is: 1.894250166s
//^^ finally all decremented and came to 0. program ends <- wg.Add(1)

// Without go routine it takes 5.253426667s to complete with go routines it takes only 1.894250166s. 