package main

import (
	"fmt"
	"time"
)

func main() {
	signal := make(chan int, 3)
	go printIt(signal)
	fmt.Println("Waiting for answer")
	answer := <-signal
	fmt.Printf("Answer: %v\n", answer)
	answer = <-signal
	fmt.Printf("Answer: %v\n", answer)

	for _ = range signal {

	}

	for {
		select {
		case x := <-signal:
			fmt.Println("Got signal", x)
			return
		default:
			fmt.Println("Booooring!")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func printIt(sdfsdf chan int) {
	time.Sleep(2 * time.Second)
	fmt.Println("Waking up")
	sdfsdf <- 4
	fmt.Println("Sending")
	sdfsdf <- 5
	fmt.Println("Waking up again")
	fmt.Println("Waiting 5 seconds")
	time.Sleep(5 * time.Second)
	sdfsdf <- 16
	close(sdfsdf)
	sdfsdf <- 25
}
