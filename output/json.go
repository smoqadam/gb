package output

import (
	"encoding/json"
	"github.com/smoqadam/gb/benchmark"
)

type Json struct{}

func (j Json) Export(m *benchmark.Metrics) ([]byte, error) {
	output, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return output, nil
}
