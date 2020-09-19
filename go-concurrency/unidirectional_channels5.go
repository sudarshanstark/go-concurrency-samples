package main // This is the corrected code for unidirectional_channels4.go

import "fmt"

func display(myChannel chan<- int) { // Converting to send only.
	fmt.Println("Display goroutine")
	myChannel <- 10
}
func main() {
	myChannel := make(chan int)
	go display(myChannel)
	number := <-myChannel
	fmt.Printf("Main Goroutine %d", number)
}
