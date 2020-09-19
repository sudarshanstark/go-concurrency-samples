package main

import (
	"fmt"
	"time"
)

func display(name string) {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
		fmt.Println(name)
	}
}
func main() {
	go display("First")
	go display("Second")
}
