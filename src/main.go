package main

import (
	"./configuration"
	"flag"
	"fmt"
)

func main() {
	configPtr := flag.String("configuration", "config.json", "a path to config file")
	config := configuration.ReadConfigurationFromFile(*configPtr)
	flag.Parse()

	Test(config)
}

func Test(config configuration.Configuration) {
	fmt.Printf("Testing url -> %s \n", config.ServerUrl)
}
