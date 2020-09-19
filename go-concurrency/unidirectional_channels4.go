package main // Creating a send only channel and declaring it to use it as bi-diectional channel but it's not possible.
// It is impossibe to convert uni-directional channel to bi-directional channel.
// The code below is wrong.
import "fmt"

func display(myChannel chan int) {
	fmt.Println("Display goroutine")
	myChannel <- 10
}
func main() {
	myChannel := make(chan<- int)
	go display(myChannel)
	number := <-myChannel
	fmt.Printf("Main Goroutine %d", number)
}
