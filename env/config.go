package env

import "fmt"

type Config struct {
	Name             string `default:"hello-world" desc:"Name of the instance" split_words:"true"`
	Namespace        string `default:"hello-ns" desc:"Namespace of the instance" split_words:"true"`
	CABundleFilePath string `desc:"Path to CABundle" split_words:"true"`
}

// func (c *Config) initialize() {
// 	c.initializeCABundle()
// }

func (c *Config) InitializeCABundle() {
	if len(c.CABundleFilePath) != 0 {
		return
	}
	fmt.Println(c.CABundleFilePath)
}
