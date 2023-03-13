package main

import (
	c "demo/concurrency"
	dsa "demo/dsa"
	config "demo/env"
	io "demo/io_opts"
	"fmt"
	"runtime"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// var remoteRouteMap map[string][]string

func main() {
	start := time.Now()
	defer func() {
		// checking time taken to run the func
		fmt.Println("Time take to execute the func is", time.Since(start))
		// fetching number of gorutines spawn by our profram
		fmt.Println("num of gorutine spawn is", runtime.NumGoroutine())
	}()

	// initiating a channel of type bolean
	smokeSignal := make(chan bool)
	sumChan := make(chan int)

	evilNinjas := "Hello"
	// Using go keyword to spawn multiple goroutines / processes
	go c.Attack(evilNinjas, smokeSignal)
	// printing the recieve signal from attack func
	fmt.Println(<-smokeSignal)

	go c.Sum(1, 2, sumChan)
	fmt.Println(<-sumChan)

	// Wait until all the threads are getting executed
	// time.Sleep(2 * time.Second)

	dsa.LinkedList()
	dsa.Stack()
	dsa.Queue()
	dsa.Bst()
	dsa.Trie()

	// SetUp envconfig
	var conf = new(config.Config)
	if err := envconfig.Usage("hello", conf); err != nil {
		panic(err)
	}
	if err := envconfig.Process("hello", conf); err != nil {
		panic(err)
	}
	// Call the func
	conf.InitializeCABundle()

	writeConfig := &io.WriteConfig{
		CloudName: "hello",
		Region:    "India",
	}
	io.WriteToFile(*writeConfig)
}
