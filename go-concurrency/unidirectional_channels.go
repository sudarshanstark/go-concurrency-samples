package main

import "fmt"

func display(myChannel chan int) {
	fmt.Println("Display goroutine")
	myChannel <- 10
}
func main() {
	myChannel := make(chan int)
	go display(myChannel)
	number := <-myChannel
	fmt.Printf("Main Goroutine %d", number)
}
