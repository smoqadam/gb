package output

import (
	"github.com/smoqadam/gb/benchmark"
)

type Exporter interface {
	Export(metrics *benchmark.Metrics) ([]byte, error)
}

func NewExporter(format string) (Exporter, error) {
	switch format {
	case "json":
		return Json{}, nil
	default:
		return Stdout{}, nil
	}
}
