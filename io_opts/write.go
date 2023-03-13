package ioopts

import (
	"demo/utils"
	"encoding/json"
	"log"
	"os"
	"path"
)

type WriteConfig struct {
	CloudName string `json:"CloudName"`
	Region    string `json:"Region"`
}

func WriteToFile(cluster WriteConfig) error {
	filePath := path.Join(utils.ProjectRoot(), "hello.json")

	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Failed to create fixture file")
		return err
	}
	defer file.Close()
	bytefile, err := json.MarshalIndent(cluster, "", " ")
	if err != nil {
		log.Println("error occurred while MarshalIndent json file", err)
		return err
	}
	err = os.WriteFile(filePath, bytefile, 0777)
	if err != nil {
		log.Println("error occurred while writing json file", err)
		return err
	}
	return nil
}
