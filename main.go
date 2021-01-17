package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	configPtr := flag.String("configuration", "config.json", "a path to config file")
	flag.Parse()
	content, err := ioutil.ReadFile(*configPtr)
	if err != nil {
		log.Fatal(err)
	}
	configStr := string(content)

	var config Configuration
	json.Unmarshal([]byte(configStr), &config)

	Test(config)
}

func Test(config Configuration){
	fmt.Printf("Testing url -> %s \n", config.ServerUrl)
}

type Configuration struct {
	ServerUrl string
}
