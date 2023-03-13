package ioopts

import (
	"fmt"
	"os"
)

var filePath = "/home/avesha/Documents/golang/prac/hello.txt"

func IoOpts() {
	r, err := os.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(r)
}
