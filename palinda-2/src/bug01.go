package main

import "fmt"

/*

	The reason for the bug is that there are no gorutines to handle the channel
	By adding a gorutine to handle the channel, the program will work as expected

	Now the ch <- "Hello world!" simply sends the message to the channel and the gorutine will handle the message and print it


*/

// I want this program to print "Hello world!", but it doesn't work.
func main() {
	ch := make(chan string)

	go func() {
		fmt.Println(<-ch)
	}()

	ch <- "Hello world!"
}
