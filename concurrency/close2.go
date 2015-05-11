package main

import "fmt"

func main() {
	c1 := sendSomeInts()

Outer:
	for {
		select {
		case i, ok := <-c1:
			if ok {
				fmt.Printf("Channel says %d\n", i)
			} else {
				break Outer
			}
		}
	}
}

func sendSomeInts() <-chan int {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	return c
}

// end-of-code - OMIT
