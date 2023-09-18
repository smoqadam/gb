package benchmark

import (
	"math"
	"sync"
	"time"
)

type Metrics struct {
	mu           sync.Mutex
	FastestTime  time.Duration
	SlowestTime  time.Duration
	AverageTime  time.Duration
	TotalTime    time.Duration
	ErrorCount   int
	SuccessCount int
	Response2xx  int // Count of 2xx responses
	Response3xx  int // Count of 3xx responses
	Response4xx  int // Count of 4xx responses
	Response5xx  int // Count of 5xx responses
}

func NewMetrics() Metrics {
	return Metrics{
		SlowestTime: 0,
		TotalTime:   0,
		FastestTime: time.Duration(math.MaxInt64),
	}
}
