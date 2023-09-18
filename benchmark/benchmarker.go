package benchmark

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"strings"
	"sync"
	"time"
)

func Start(config Config) {
	fmt.Println("Start benchmarking: ", config)
	r := rate.NewLimiter(rate.Every(time.Second), config.Limit)
	durations := make(chan time.Duration, config.Number)
	errs := make(chan error, config.Concurrent)
	semaphore := make(chan struct{}, config.Concurrent) // to control how many concurrent request can be run

	var wg sync.WaitGroup
	wg.Add(config.Number)

	m := NewMetrics()
	for i := 1; i <= config.Number; i++ {
		if err := r.Wait(context.Background()); err != nil {
			continue
		}

		// program waits if number of tokens in semaphore is at its maximum capacity,
		// otherwise acquire a token and continue
		semaphore <- struct{}{}
		go func() {
			executeRequest(config, durations, errs, &m, &wg)
			<-semaphore // release a token
		}()
	}

	wg.Wait()

	// Why closing channels before the end of the function?
	// Since all the goroutines are done by the time we reach that point,
	// we know that no more values will be sent on these channels, so it's safe to close them.
	//
	// After closing them, you can safely use the range keyword to iterate over the values
	// stored in these channels. If you don't close the channels, using range to iterate over
	// them will cause the program to be stuck in an infinite loop because range will keep waiting
	// for more values to come.
	close(durations)
	close(errs)

	for d := range durations {
		if d < m.FastestTime {
			m.FastestTime = d
		}
		if d > m.SlowestTime {
			m.SlowestTime = d
		}
		m.TotalTime += d
		m.SuccessCount++
	}
	m.AverageTime = m.TotalTime / time.Duration(config.Number)
	for range errs {
		m.ErrorCount++
	}

	fmt.Printf("Average Time: %v\n", m.AverageTime)
	fmt.Printf("Total Time: %v\n", m.TotalTime)
	fmt.Printf("Fastest Time: %v\n", m.FastestTime)
	fmt.Printf("Slowest Time: %v\n", m.SlowestTime)
	fmt.Printf("Error Count: %d\n", m.ErrorCount)
	fmt.Printf("Success Count: %d\n", m.SuccessCount)
	fmt.Printf("200: %d\n", m.Response2xx)
	fmt.Printf("3xx: %d\n", m.Response3xx)
	fmt.Printf("4xx: %d\n", m.Response4xx)
	fmt.Printf("5xx: %d\n", m.Response5xx)
	fmt.Printf("Bytes Received: %db\n", m.ContentLength)
}

func executeRequest(c Config, durations chan<- time.Duration, errs chan<- error, metrics *Metrics, wg *sync.WaitGroup) {
	defer wg.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.Timeout))
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", c.URL, nil)

	for _, h := range c.Headers {
		key, value, err := parseHeader(h)
		if err != nil {
			errs <- err
			continue
		}
		req.Header.Set(key, value)
	}

	start := time.Now()
	res, err := http.DefaultClient.Do(req)
	elapsed := time.Since(start)
	if err != nil {
		errs <- err
		return
	}

	defer res.Body.Close()
	durations <- elapsed
	metrics.Update(res)
}

func parseHeader(header string) (key, value string, err error) {
	parts := strings.SplitN(header, ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid header format: %s", header)
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), nil
}
