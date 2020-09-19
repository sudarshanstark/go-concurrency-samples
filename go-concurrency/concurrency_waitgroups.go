package main

import (
	"fmt"
	"sync"
	"time"
)

// Waitgrous are group of values which are used to control/maintain the state of goroutine calls.
// WaitGroups will wait untill a collection of goroutines completes their work.
// Wait groups will block the main routine from execution untill all the go routines finishes their work.
func dispay(i int, g *sync.WaitGroup) {
	fmt.Println("Start", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End", i)
	g.Done() // Statement to specify the waitgroup about the completion of a goroutine.
}
func main() {
	var g sync.WaitGroup
	for i := 1; i <= 5; i++ {
		g.Add(1) // This statement creates a waitgroup entry for each goroutine invocation.Therefore it create 5 waitgroup entries.
		go dispay(i, &g)
	}
	g.Wait() // This is a blocking call to ensure goroutine waits until all goroutine gets completed.
	fmt.Println("Exiting Main Goroutine")
}
