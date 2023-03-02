/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"the-go-programming-language/internal/functions/functionvalues"
)

var c5e13SearchTerms []string

// c5e13Cmd represents the c5e13 command
var c5e13Cmd = &cobra.Command{
	Use:   "c5e13",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		functionvalues.RunCrawl(c5e13SearchTerms)
	},
}

func init() {
	rootCmd.AddCommand(c5e13Cmd)
	searchCmd.Flags().StringSliceVarP(&c5e13SearchTerms, "work-list", "w", []string{"https://golang.org"}, "An comma separated string of crawl targets")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c5e13Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c5e13Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
