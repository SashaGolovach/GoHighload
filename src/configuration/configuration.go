package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	ServerUrl string
}

func ReadFromFile(path string) Configuration {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	configStr := string(content)

	var config Configuration
	json.Unmarshal([]byte(configStr), &config)

	return config
}
