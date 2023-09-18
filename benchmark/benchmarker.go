package benchmark

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
	"time"
)

func Start(url string, n int, c int, limit int) {

	fmt.Println("Start benchmarking: ", url)
	r := rate.NewLimiter(rate.Every(time.Second), limit)
	durations := make(chan time.Duration, n)
	errs := make(chan error, c)
	semaphore := make(chan struct{}, c) // to control how many concurrent request can be run

	var wg sync.WaitGroup
	wg.Add(n)

	m := NewMetrics()
	for i := 1; i <= n; i++ {
		if err := r.Wait(context.Background()); err != nil {
			continue
		}

		// program waits if number of tokens in semaphore is at its maximum capacity,
		// otherwise acquire a token and continue
		semaphore <- struct{}{}
		go func() {
			test(url, durations, errs, &wg)
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
	m.AverageTime = m.TotalTime / time.Duration(n)
	for range errs {
		m.ErrorCount++
	}

	fmt.Printf("Average Time: %v\n", m.AverageTime)
	fmt.Printf("Total Time: %v\n", m.TotalTime)
	fmt.Printf("Fastest Time: %v\n", m.FastestTime)
	fmt.Printf("Slowest Time: %v\n", m.SlowestTime)
	fmt.Printf("Error Count: %d\n", m.ErrorCount)
	fmt.Printf("Success Count: %d\n", m.SuccessCount)
}

func test(url string, durations chan<- time.Duration, errs chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	res, err := http.Get(url)
	elapsed := time.Since(start)
	if err != nil {
		errs <- err
		return
	}
	defer res.Body.Close()
	durations <- elapsed
}
