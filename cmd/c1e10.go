/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"the-go-programming-language/internal/tutorial"

	"github.com/spf13/cobra"
)

var c1e10urls []string

// c1e10Cmd represents the c1e10 command
var c1e10Cmd = &cobra.Command{
	Use:   "c1e10",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tutorial.PrintFetchAll(c1e10urls)
		tutorial.PrintFetchAll(c1e10urls)
	},
}

func init() {
	rootCmd.AddCommand(c1e10Cmd)
	c1e10Cmd.Flags().StringSliceVarP(&c1e10urls, "urls", "u", []string{"https://maxsargentdev.github.io/"}, "An comma separated string of URLs")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c1e10Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c1e10Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// First iteration much slower
// 0.21s    3968 https://maxsargentdev.github.io/
// 0.21s elapsed
// 0.01s    3968 https://maxsargentdev.github.io/
// 0.02s elapsed
