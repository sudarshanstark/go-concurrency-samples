package main // Recieve only channels. We can't write value to the channel, we can just recieve or read data from it.
// The code below is wrong.
import "fmt"

func display(myChannel <-chan int) {
	fmt.Println("Display goroutine")
	// myChannel <- 10
}
func main() {
	myChannel := make(<-chan int) // Read Only Channels.
	go display(myChannel)
	number := <-myChannel // Reading data from channel to a variable.
	fmt.Printf("Main Goroutine %d", number)
}
