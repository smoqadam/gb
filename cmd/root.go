package cmd

import (
	"fmt"
	"github.com/smoqadam/gb/benchmark"
	"os"

	"github.com/spf13/cobra"
)

var (
	number     int
	concurrent int
	limit      int
	url        string
	timeout    int
	headers    []string
	rootCmd    = &cobra.Command{
		Use:   "gb",
		Short: "A benchmarking tool",
		Long:  `An experimental HTTP benchmarking tool written in Go`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("N:", number)
			config := benchmark.Config{
				Concurrent: concurrent,
				Number:     number,
				Limit:      limit,
				URL:        url,
				Timeout:    timeout,
				Headers:    headers,
			}
			benchmark.Start(config)
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

	rootCmd.Flags().IntVarP(&number, "number", "n", 10, "number of request")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 10, "limit of concurrent requests per second")
	rootCmd.Flags().IntVarP(&concurrent, "concurrent", "c", 10, "number of concurrent request")
	rootCmd.Flags().IntVarP(&timeout, "timeout", "T", 30, "Timeout for each requests (in seconds)")

	rootCmd.Flags().StringSliceVarP(&headers, "header", "H", nil, "Header to pass to request. This flag can be used multiple times")
}
