/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/composite-types/slices"

	"github.com/spf13/cobra"
)

var c4e5inputSlice []string

// c4e5Cmd represents the c4e5 command
var c4e5Cmd = &cobra.Command{
	Use:   "c4e5",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		slices.RemoveAdjacentDupes(c4e5inputSlice)
	},
}

func init() {
	rootCmd.AddCommand(c4e5Cmd)
	c4e5Cmd.Flags().StringSliceVarP(&c4e5inputSlice, "input", "i", []string{"one", "two", "two", "three", "two", "three", "three", "three", "two", "four"}, "Input string slice for dedupe operation")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c4e5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c4e5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
