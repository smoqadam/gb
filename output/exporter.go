package output

import (
	"errors"
	"github.com/smoqadam/gb/benchmark"
)

type Exporter interface {
	Export(metrics *benchmark.Metrics) ([]byte, error)
}

func NewExporter(format string) (Exporter, error) {
	switch format {
	case "json":
		return Json{}, nil
	case "cli":
		return Cli{}, nil
	default:
		return nil, errors.New("unsupported export format: " + format)
	}
}
