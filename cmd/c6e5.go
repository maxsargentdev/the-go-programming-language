/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/methods/intset2"

	"github.com/spf13/cobra"
)

var c6e5IntSet intset2.IntSet

// c6e5Cmd represents the c6e5 command
var c6e5Cmd = &cobra.Command{
	Use:   "c6e5",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		testIntSet := intset2.IntSet{}
		testIntSet.Add(1)
		testIntSet.Add(2)
		testIntSet.Add(3)
		testIntSet.Add(4)
		testIntSet.Add(5)
		testIntSet.SymmetricDifference(&c6e5IntSet)
		println(testIntSet.String())
	},
}

func init() {
	rootCmd.AddCommand(c6e5Cmd)
	c6e5IntSet.Add(1)
	c6e5IntSet.Add(2)
	c6e5IntSet.Add(3)
	c6e5IntSet.Add(4)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c6e5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c6e5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
