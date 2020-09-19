package main

import "fmt" // The code below is wrong.
// Comment the line no.9 to make it right.
func main() {
	channel := make(chan int, 2) // Create a buffered channe of capacity 2
	channel <- 10                //Sending three values to the channel.
	channel <- 20
	channel <- 30 // The capacity of channel is exhausted.
	// You cannot write 3 values to channel which only accepts 2 values.
	// Once the buffer slots gets filled up, then the send channels will be blocking as usual.
	fmt.Printf("%d\n", <-channel)
	fmt.Printf("%d", <-channel)
}
