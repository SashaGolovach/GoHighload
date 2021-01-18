package main

import (
	"./configuration"
	"./loadTesting"
	"encoding/json"
	"flag"
	"io/ioutil"
)

func main() {
	configPtr := flag.String("configuration", "config.json", "a path to config file")
	config := configuration.ReadFromFile(*configPtr)
	flag.Parse()
	report := loadTesting.Run(config)
	bytes, err := json.Marshal(report)
	if err == nil{
		ioutil.WriteFile("./report.json", bytes, 0644)
	}
}
