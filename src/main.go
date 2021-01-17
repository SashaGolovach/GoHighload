package main

import (
	"./configuration"
	"./loadTesting"
	"flag"
)

func main() {
	configPtr := flag.String("configuration", "config.json", "a path to config file")
	config := configuration.ReadFromFile(*configPtr)
	flag.Parse()

	loadTesting.Run(config)
}
