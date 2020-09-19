package main

import "fmt"

func display(myChannel chan int) {
	fmt.Println("Display Goroutine")
	for i := 1; i <= 5; i++ {
		myChannel <- i // Send set of i values iteratively to the channel myChannel.
	}
	close(myChannel)
}
func main() {
	myChannel := make(chan int)
	go display(myChannel)
	for {
		num, status := <-myChannel // Read a value from myChannel.
		if status == false {       // If channel is closed.
			break
		}
		fmt.Printf("Main Goroutine:%d\n", num)
	}
}
