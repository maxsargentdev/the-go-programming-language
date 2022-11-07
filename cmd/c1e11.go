/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"the-go-programming-language/internal/tutorial"

	"github.com/spf13/cobra"
)

var c1e11urls []string

// c1e11Cmd represents the c1e11 command
var c1e11Cmd = &cobra.Command{
	Use:   "c1e11",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tutorial.PrintFetchAll(c1e11urls) // provide list via flag with default, use long list
	},
}

func init() {
	rootCmd.AddCommand(c1e11Cmd)
	c1e11Cmd.Flags().StringSliceVarP(&c1e11urls, "urls", "u", []string{"https://maxsargentdev.github.io/", "https://google.com", "https://amazon.com"}, "An comma separated string of URLs") // need a much longer list to go here
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c1e11Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c1e11Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// 0.07s    3968 https://maxsargentdev.github.io/
// 0.17s   15070 https://google.com
// 0.50s    6591 https://amazon.com
// 0.50s elapsed
