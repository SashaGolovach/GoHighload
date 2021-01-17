package loadTesting

import (
	"../configuration"
	"fmt"
	"net/http"
)

func Run(config configuration.Configuration) {
	fmt.Printf("Testing url -> %s \n", config.ServerUrl)

	resp, err := http.Get(config.ServerUrl)
	if err != nil {
		println("Got error -> ", err.Error())
	} else {
		println("Got response -> ", resp.Status)
	}
}
