package main

import (
	"context"
	"demo/concepts"
	c "demo/concurrency"
	dsa "demo/dsa"
	config "demo/env"
	io "demo/io_opts"
	"demo/monitoing"
	"fmt"
	"runtime"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/common/model"
)

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

	// calling networkinterface client to get interface info
	m := concepts.NewNetInfClientProvider()
	info, _ := m.GetNetworkInterfaceInfo("tcp:192.3200")
	fmt.Println(info)

	// encoding yaml data storing into a file and reading and printing the data
	// after decoding using yaml.unmarshel().
	concepts.YamlFunc()

	// Sum is a vaiadic function which can accept n number of arguments
	totalSum := concepts.Sum(1, 2, 3, 4, 5, 6)
	fmt.Println(totalSum)

	concepts.GetScheduleTime()

	// Initializing chan of type bool
	// done := make(chan bool)

	// fmt.Println("------------------ Starting Healthcheck -------------------")
	// cfg := &c.Config{}
	// cfg.PeriodicHealthCheck(done)
	// fmt.Println("------------------ Ending Healthcheck ---------------------")

	// start node health monitoring

	res, err := monitoing.PrometheusQuery(context.Background(), "kube_statefulset_created")
	if err != nil {
		fmt.Println(err, "error during executing prometheus query")
		panic(err)
	}
	vector := res.(model.Vector)
	if vector.Len() == 0 {
		fmt.Println("This query returned no data")
	}
	fmt.Println(vector, "<<<<<<<<<<< value from prometheus metrics query")

	// end node health monitoring

	// server mode

	// monitoing.StartMetricsCollector("8081")

	// end server mode

}
