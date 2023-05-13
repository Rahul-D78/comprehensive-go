package concepts

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// YAML tags are used to specify how Go struct fields should be encoded as YAML data.
type Person struct {
	Name string `yaml:"name,omitempty"`
	Age  int    `yaml:"age"`
}

func YamlFunc() error {
	person := &Person{
		Name: "hello",
		Age:  10,
	}

	y, err := yaml.Marshal(person)
	if err != nil {
		return err
	}
	fmt.Println("Yaml marshled data", y)

	err = os.WriteFile("config.yaml", y, 0644)
	if err != nil {
		return err
	}

	source, err := os.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	p := &Person{}
	err = yaml.Unmarshal(source, p)
	if err != nil {
		return err
	}
	fmt.Println(p.Name)

	return nil
}
