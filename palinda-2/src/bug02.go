package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 The problem with the code is that the program exits before the goroutine has finished.
 By adding a wait group, the main function will wait for the goroutine to finish before exiting.
 And with that it wont skip the last number and print out the entire series.


*/

// This program should go to 11, but it seemingly only prints 1 to 10.
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup // create wait group
	wg.Add(1)             // add one to the counter
	go Print(ch, &wg)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait() // wait for the goroutine to finish
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wg *sync.WaitGroup) {
	for n := range ch { // reads from channel until it's closed
		time.Sleep(10 * time.Millisecond) // simulate processing time
		fmt.Println(n)
	}
	wg.Done() // remove one from the counter
}
