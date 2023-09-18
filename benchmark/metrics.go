package benchmark

import (
	"io"
	"math"
	"net/http"
	"sync"
	"time"
)

type Metrics struct {
	mu            sync.Mutex
	FastestTime   time.Duration `json:"fastest_time"`
	SlowestTime   time.Duration `json:"slowest_time"`
	AverageTime   time.Duration `json:"average_time"`
	TotalTime     time.Duration `json:"total_time"`
	ErrorCount    int           `json:"error_count"`
	SuccessCount  int           `json:"success_count"`
	Response2xx   int           `json:"response_2_xx"` // Count of 2xx responses
	Response3xx   int           `json:"response_3_xx"` // Count of 3xx responses
	Response4xx   int           `json:"response_4_xx"` // Count of 4xx responses
	Response5xx   int           `json:"response_5_xx"` // Count of 5xx responses
	ContentLength int64         `json:"content_length"`
}

func NewMetrics() Metrics {
	return Metrics{
		SlowestTime: 0,
		TotalTime:   0,
		FastestTime: time.Duration(math.MaxInt64),
	}
}

func (m *Metrics) Update(res *http.Response) {

	m.mu.Lock()

	b, _ := io.ReadAll(res.Body)
	m.ContentLength += int64(len(b))

	switch {
	case res.StatusCode >= 200 && res.StatusCode < 300:
		m.Response2xx++
	case res.StatusCode >= 300 && res.StatusCode < 400:
		m.Response3xx++
	case res.StatusCode >= 400 && res.StatusCode < 500:
		m.Response4xx++
	case res.StatusCode >= 500:
		m.Response5xx++
	}
	m.mu.Unlock()
}
