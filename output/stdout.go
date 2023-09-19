package output

import (
	"fmt"
	"github.com/smoqadam/gb/benchmark"
)

type Stdout struct {
}

func (c Stdout) Export(m *benchmark.Metrics) ([]byte, error) {

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

	return nil, nil
}
