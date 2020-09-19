package main

import "fmt"

func main() {
	// There is no wait/sleep operation in main goroutine.
	channel := make(chan int, 2) // Create a buffered channe of capacity 2
	channel <- 10                //Sending two values to the channel.
	channel <- 20
	fmt.Printf("%d\n", len(channel)) // Length of buffered channel is 2
	// Length will change based on the current status of the buffer.
	fmt.Printf("%d\n", <-channel) // After reading a value from the channel the buffer value is 1
	fmt.Printf("%d\n", len(channel))
	fmt.Printf("%d\n", <-channel)  //Recieving two values from the channel.
	fmt.Printf("%d", cap(channel)) // Prints the capacity of buffer.
	// How many times if you have read/write the capacity remains the same.
}
