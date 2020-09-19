package main

import (
	"fmt"
	"time"
)

func wish() {
	fmt.Println("Wish Goroutine")
}
func main() {
	go wish()
	time.Sleep(1 * time.Second)
	fmt.Println("Inside main")
}
