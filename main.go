package main

import (
	"fmt"
	"time"
)

func attack(target string) {
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(1 * time.Second)
}

func main() {
	// checking time taken to run the func
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evilNinjas := []string{"Nano", "Pico", "Fermi", "Auto"}
	for _, evilNinja := range evilNinjas {
		// Using go keyword to spawn multiple goroutines / processes
		go attack(evilNinja)
	}

	// For temp purpose we can wait for 2 sec to exit out of main process
	time.Sleep(2 * time.Second)
}
