package benchmark

import (
	"math"
	"time"
)

type Metrics struct {
	FastestTime  time.Duration
	SlowestTime  time.Duration
	AverageTime  time.Duration
	TotalTime    time.Duration
	ErrorCount   int
	SuccessCount int
}

func NewMetrics() Metrics {
	return Metrics{
		SlowestTime: 0,
		TotalTime:   0,
		FastestTime: time.Duration(math.MaxInt64),
	}
}
