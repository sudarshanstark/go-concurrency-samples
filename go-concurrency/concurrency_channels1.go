package main

import "fmt"

//If there is a read statement given then always there should be a write statement to complement read statement.
// If there is a read operation from a channel inside a go routine, then there should be a proper write operation to supply data to the channel
// Otherwise, go compiler will raise a panic of deadlock.
func display(myChannel chan int) {
	fmt.Println("Display goroutine")
	myChannel <- 10 // Writing the value to the int channel myChannel.
}
func main() {
	myChannel := make(chan int)
	go display(myChannel)
	num := <-myChannel // Reading data from channel myChannel.
	fmt.Println(num)
	// Whenever a channel read is happening the go compiler keeps flow blocked until a write is being happening on the channel.
	fmt.Println("Inside main")
}
