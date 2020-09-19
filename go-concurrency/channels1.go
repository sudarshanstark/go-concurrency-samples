package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	process1 := <-doit(1, 1, &wg)
	process2 := <-doit(2, 2, &wg)
	process3 := <-doit(3, 3, &wg)

	fmt.Println(process1)
	fmt.Println(process2)
	fmt.Println(process3)

	wg.Wait()
}

func doit(input_id int, delay_time time.Duration, wg *sync.WaitGroup) chan string {
	output := make(chan string)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * delay_time)
		output <- fmt.Sprintf("Process %d done", input_id)
	}()

	return output
}
