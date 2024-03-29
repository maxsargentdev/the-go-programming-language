/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"the-go-programming-language/internal/methods/intset"

	"github.com/spf13/cobra"
)

var c6e4IntSet intset.IntSet

// c6e4Cmd represents the c6e4 command
var c6e4Cmd = &cobra.Command{
	Use:   "c6e4",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		for k, v := range c6e4IntSet.Elems() {
			fmt.Printf("%d - %d\n", k, v)
		}
	},
}

func init() {
	rootCmd.AddCommand(c6e4Cmd)

	c6e4IntSet.Add(1)
	c6e4IntSet.Add(2)
	c6e4IntSet.Add(3)
	c6e4IntSet.Add(4)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c6e4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c6e4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
