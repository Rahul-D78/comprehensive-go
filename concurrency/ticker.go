package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	once sync.Once
}

func (c *Config) PeriodicHealthCheck(done chan bool) {
	c.once.Do(func() {
		go healthChek(done)
	})
}

func healthChek(done chan bool) {
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			err := monitor()
			if err != nil {
				fmt.Println("error monitoring system")
			}
		case <-done:
			fmt.Println("health check monitor done")
			ticker.Stop()
			return
		}
	}
}

func monitor() error {
	fmt.Println("monitoring health of system")
	return nil
}
