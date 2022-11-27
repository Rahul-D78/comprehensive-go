package concurrency

import (
	"fmt"
	"time"
)

func Attack(target string, attacked chan bool) {
	fmt.Println("Throwing ninja stars at", target)
	time.Sleep(1 * time.Second)
	// after our work done we can send a signal back to the main func
	attacked <- true
}

func Sum(a, b int, chansig chan int) {
	res := a + b
	chansig <- res
}
