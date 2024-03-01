// Stefan Nilsson 2013-03-13

// This is a testbed to help you understand channels better.
package main

/*
	If I increase the number of consumers the program would run faster as there are more consumers to handle the data being pushed to the channel.

	If Wait and close change places, the program will throw an err beacuse the channel will be closed before the producers are done.

	If I were to remove the close(ch), the program would still exit fine, but in a larger program having open channels may lead to performance cost, it's always best practise to clean up the channels.

	Because I wait for all the producers to finish but not consumers before closing the channel, there may be consumers who don't get to print all the strings, and so I cant be certain all of the consumers finish printing their strings.

	If I move the close(ch) to Produce, the program will throw an err beacuse the channel will be closed before all the producers are done, only one of the producers will be able to finish.


*/

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 2

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup)
	wgp.Add(producers)
	wgc := new(sync.WaitGroup)
	wgc.Add(consumers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch, wgc)
	}
	// using this go rutine to close the channel after all producers are done, but then allows wgc to continue until it's all done reading the inputs
	go func() {
		wgp.Wait() // Wait for all producers to finish.
		close(ch)
	}()
	wgc.Wait() // Wait for all consumers to finish.

	fmt.Println("time:", time.Now().Sub(before))
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string, wg *sync.WaitGroup) {
	for s := range ch {
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
	wg.Done()
}

// RandomSleep waits for x ms, where x is a random number, 0 < x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}
