package concepts

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type GitProperties struct {
	CommitPrefix  string `json:"commitPrefix,omitempty"`
	PrDescription string `json:"prDescriptionFormat,omitempty"`
}

func YamlFunc() error {
	gp := &GitProperties{
		CommitPrefix:  "abc",
		PrDescription: "efg",
	}

	g, err := yaml.Marshal(gp)
	if err != nil {
		return err
	}
	fmt.Println("Yaml marshled data for git properties", string(g))

	err = os.WriteFile("gp.yaml", g, 0644)
	if err != nil {
		return err
	}

	source, err := os.Open("gp.yaml")
	if err != nil {
		return err
	}

	var lines []string

	scanner := bufio.NewScanner(source)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fmt.Println(lines)

	// Parse YAML data into struct
	var gitp GitProperties
	yamlData := strings.Join(lines, "\n")
	err = yaml.Unmarshal([]byte(yamlData), &gitp)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v", err)
	}

	fmt.Println()
	fmt.Println("Server:", gitp.CommitPrefix)

	return nil
}
