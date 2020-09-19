// Channels will pass the data to the Go-routines.
package main

import "fmt"

func main() {
	var myChannel chan int
	if myChannel == nil {
		fmt.Println("Channel not initialized.")
	}
	myChannel = make(chan int)
	fmt.Println(myChannel)
	// data := <- myChannel // Read from the channel.
	// myChannel <- data // Writing to the channel.
}
