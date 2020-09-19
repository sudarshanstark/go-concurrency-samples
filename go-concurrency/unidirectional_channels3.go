package main // Converting Bi-directional channel from line no.9 to Uni-directional channel in line no.4
import "fmt" // It is possible to convert Bi-directional channel to either send only or recieve only.
func display(myChannel chan<- int) { // Converting the channel to Send only.
	fmt.Println("Display goroutine")
	myChannel <- 10
}
func main() {
	myChannel := make(chan int) //Initializing the channel to bi-directional.
	go display(myChannel)
	number := <-myChannel
	fmt.Printf("Main Goroutine %d", number)
}
