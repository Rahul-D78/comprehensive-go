package main

import (
	"fmt"
	"time"
)

func attack(target string, attacked chan bool) {
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(1 * time.Second)
	// after our work done we can send a signal back to the main func
	attacked <- true
}

func main() {
	// checking time taken to run the func
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	smokeSignal := make(chan bool)

	evilNinja := "Nano"
	// Using go keyword to spawn multiple goroutines / processes
	go attack(evilNinja, smokeSignal)
	// printing the recieve signal from attack func
	fmt.Println(<-smokeSignal)

	// For temp purpose we can wait for 2 sec to exit out of main process
	// time.Sleep(2 * time.Second)
}
