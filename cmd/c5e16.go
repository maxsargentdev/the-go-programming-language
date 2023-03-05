/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/functions/variadic"

	"github.com/spf13/cobra"
)

var c5e16Sep string
var c5e16Strings []string

// c5e16Cmd represents the c5e16 command
var c5e16Cmd = &cobra.Command{
	Use:   "c5e16",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		variadic.Join(c5e16Sep, c5e16Strings...)
	},
}

func init() {
	rootCmd.AddCommand(c5e16Cmd)
	c5e16Cmd.Flags().StringVarP(&c5e16Sep, "separator", "s", ",", "A separator for the join function")
	c5e16Cmd.Flags().StringSliceVarP(&c5e16Strings, "terms", "t", []string{"One", "Two", "Three", "Four", "Five"}, "An comma separated string of integer search terms")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c5e16Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c5e16Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
