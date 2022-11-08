/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"the-go-programming-language/internal/tutorial"

	"github.com/spf13/cobra"
)

var c1e9urls []string

// c1e9Cmd represents the c1e9 command
var c1e9Cmd = &cobra.Command{
	Use:   "c1e9",
	Short: "Chapter 1 Exercise 9 - Fetch with status code",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("c1e9 called")
		tutorial.StatusCodePrefixCheckFetch(c1e9urls)
	},
}

func init() {
	rootCmd.AddCommand(c1e9Cmd)
	c1e9Cmd.Flags().StringSliceVarP(&c1e9urls, "urls", "u", []string{"www.google.com"}, "An comma separated string of URLs")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c1e9Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c1e9Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
