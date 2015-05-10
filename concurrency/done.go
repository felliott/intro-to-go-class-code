package main

import "fmt"

func main() {
	done := make(chan bool)
	c := sendSomeInts(done)

Outer:
	for {
		select {
		case i := <-c:
			fmt.Printf("Channel says %d\n", i)
		case <-done:
			break Outer
		}
	}
}

func sendSomeInts(done chan<- bool) <-chan int {
	c := make(chan int)
	go func() {
		c <- 42
		done <- true
	}()
	return c
}

// end-of-code - OMIT
