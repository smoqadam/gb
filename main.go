package main

import (
	"flag"
	"fmt"
	"math"
	"net/http"
	"time"
)

func main() {
	var url string
	var n int
	flag.StringVar(&url, "url", "", "-url [URL]")
	flag.IntVar(&n, "number", 1, "-number [number of requests]")

	flag.Parse()

	start(url, n)
}

func start(url string, n int) {

	m := Metrics{
		SlowestTime: 0,
		TotalTime:   0,
		FastestTime: time.Duration(math.MaxInt64),
	}
	for i := 1; i <= n; i++ {
		d, err := test(url)
		if err != nil {
			m.ErrorCount += 1
			continue
		} else {
			m.SuccessCount += 1
		}
		if d < m.FastestTime {
			m.FastestTime = d
		}

		if d > m.SlowestTime {
			m.SlowestTime = d
		}

		m.TotalTime += d
	}
	if m.SuccessCount > 0 {
		m.AverageTime = m.TotalTime / time.Duration(m.SuccessCount)
	}
	fmt.Printf("Average Time: %v\n", m.AverageTime)
	fmt.Printf("Total Time: %v\n", m.TotalTime)
	fmt.Printf("Fastest Time: %v\n", m.FastestTime)
	fmt.Printf("Slowest Time: %v\n", m.SlowestTime)
	fmt.Printf("Error Count: %d\n", m.ErrorCount)
}

func test(url string) (time.Duration, error) {
	fmt.Println("Start fetching: ", url)
	start := time.Now()
	res, err := http.Get(url)
	elapsed := time.Since(start)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	return elapsed, nil
}

type Metrics struct {
	FastestTime  time.Duration
	SlowestTime  time.Duration
	AverageTime  time.Duration
	TotalTime    time.Duration
	ErrorCount   int
	SuccessCount int
}
