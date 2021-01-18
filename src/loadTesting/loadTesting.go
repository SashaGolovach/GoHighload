package loadTesting

import (
	"../configuration"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const ATTEMPTS_COUNT = 100_000

type Attempt struct {
	IsSuccess bool
	Error     string
}

type Report struct {
	TotalTime time.Duration
	Attempts  []Attempt
}

func Run(config configuration.Configuration) Report {
	fmt.Printf("Testing url -> %s \n", config.ServerUrl)
	var attempts [ATTEMPTS_COUNT]Attempt
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < ATTEMPTS_COUNT; i++ {
		wg.Add(1)
		go SendRequest(config.ServerUrl, &attempts[i], &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	print("\n")
	log.Printf("Processed %d requests. Elapsed time %s", ATTEMPTS_COUNT, elapsed)
	var report Report
	report.Attempts = attempts[:]
	report.TotalTime = elapsed
	return report
}

func SendRequest(url string, result *Attempt, wg *sync.WaitGroup) {
	defer wg.Done()
	defer print(".")
	_, err := http.Get(url)
	if err != nil {
		result.Error = err.Error()
	}
	result.IsSuccess = err == nil
}
