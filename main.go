package main

import (
	"fmt"
	"runtime"
	"time"
)

func captainElect(ninja chan string, message string) {
	time.Sleep(3 * time.Second)
	ninja <- message
}

func attack(target string, attacked chan bool) {
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(1 * time.Second)
	// after our work done we can send a signal back to the main func
	attacked <- true
}

func main() {
	start := time.Now()
	defer func() {
		// checking time taken to run the func
		fmt.Println("Time take to execute the func is", time.Since(start))
		// fetching number of gorutines spawn by our profram
		fmt.Println("num of gorutine spawn is", runtime.NumGoroutine())
	}()

	smokeSignal := make(chan bool)

	// initiate 2 channels in the memory
	ninja1, ninja2 := make(chan string), make(chan string)

	go captainElect(ninja1, "ninja1")
	go captainElect(ninja2, "ninja2")

	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	}

	evilNinja := "Nano"
	// Using go keyword to spawn multiple goroutines / processes
	go attack(evilNinja, smokeSignal)
	// printing the recieve signal from attack func
	fmt.Println(<-smokeSignal)

	// For temp purpose we can wait for 2 sec to exit out of main process
	// time.Sleep(2 * time.Second)

	// Implementing buffer channels (By default the capacity of a channel is 0)
	channel := make(chan string, 1)
	channel <- "Pulling val from buffer"
	fmt.Println(<-channel)
}
