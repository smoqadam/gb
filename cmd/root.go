package cmd

import (
	"fmt"
	"github.com/smoqadam/gb/benchmark"
	"github.com/smoqadam/gb/output"
	"os"

	"github.com/spf13/cobra"
)

var (
	number       int
	concurrent   int
	limit        int
	url          string
	method       string
	data         string
	outputFormat string
	timeout      int
	headers      []string
	rootCmd      = &cobra.Command{
		Use:   "gb",
		Short: "A benchmarking tool",
		Long:  `An experimental HTTP benchmarking tool written in Go`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("N:", number)
			config := benchmark.Config{
				Concurrent:  concurrent,
				Number:      number,
				Limit:       limit,
				URL:         url,
				Method:      method,
				Timeout:     timeout,
				Headers:     headers,
				RequestBody: []byte(data),
			}

			metrics := benchmark.Start(config)
			exporter, err := output.NewExporter(outputFormat)
			if err != nil {
				panic(err)
			}

			b, err := exporter.Export(metrics)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(b))
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL to benchmark")
	rootCmd.MarkFlagRequired("url")

	rootCmd.Flags().StringVarP(&outputFormat, "output", "O", "std", "output format [json, html, csv]")

	rootCmd.Flags().StringVarP(&method, "method", "m", "GET", "request method [GET, POST, PUT, PATCH, DELETE]")
	rootCmd.Flags().StringVarP(&data, "data", "d", "", "request body")
	rootCmd.Flags().IntVarP(&number, "number", "n", 10, "number of request")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 10, "limit of concurrent requests per second")
	rootCmd.Flags().IntVarP(&concurrent, "concurrent", "c", 10, "number of concurrent request")
	rootCmd.Flags().IntVarP(&timeout, "timeout", "T", 30, "Timeout for each requests (in seconds)")

	rootCmd.Flags().StringSliceVarP(&headers, "header", "H", nil, "Header to pass to request. This flag can be used multiple times")
}
