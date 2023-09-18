/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/smoqadam/gb/benchmark"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	number     int
	concurrent int
	limit      int
	url        string
	rootCmd    = &cobra.Command{
		Use:   "gb",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("N:", number)
			benchmark.Start(url, number, concurrent, limit)
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
	rootCmd.Flags().IntVarP(&concurrent, "concurrent", "c", 1, "number of concurrent request")
}
