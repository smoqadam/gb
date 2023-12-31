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
	toFile       string
	timeout      int
	headers      []string
	rootCmd      = &cobra.Command{
		Use:   "gb",
		Short: "A benchmarking tool",
		Long:  `An experimental HTTP benchmarking tool written in Go`,
		Run: func(cmd *cobra.Command, args []string) {
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
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}

			b, err := exporter.Export(metrics)
			if err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}

			if len(toFile) > 0 {
				err := os.WriteFile(toFile, b, 0666)
				if err != nil {
					fmt.Fprintf(os.Stderr, err.Error())
					os.Exit(1)
				}
				return
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

	rootCmd.Flags().StringVarP(&outputFormat, "output-format", "o", "stdout", "output format [json, html, csv]")
	rootCmd.Flags().StringVarP(&toFile, "output-file", "O", "gb", "output filename format [json, html, csv]")

	rootCmd.Flags().StringVarP(&method, "method", "m", "GET", "request method [GET, POST, PUT, PATCH, DELETE]")
	rootCmd.Flags().StringVarP(&data, "data", "d", "", "request body")
	rootCmd.Flags().IntVarP(&number, "number", "n", 10, "number of request")
	rootCmd.Flags().IntVarP(&limit, "limit", "l", 10, "limit of concurrent requests per second")
	rootCmd.Flags().IntVarP(&concurrent, "concurrent", "c", 10, "number of concurrent request")
	rootCmd.Flags().IntVarP(&timeout, "timeout", "T", 30, "Timeout for each requests (in seconds)")

	rootCmd.Flags().StringSliceVarP(&headers, "header", "H", nil, "Header to pass to request. This flag can be used multiple times")
}
