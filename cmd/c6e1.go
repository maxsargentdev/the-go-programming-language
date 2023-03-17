/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"the-go-programming-language/internal/methods/intset"

	"github.com/spf13/cobra"
)

var c6e1IntSet intset.IntSet

// c6e1Cmd represents the c6e1 command
var c6e1Cmd = &cobra.Command{
	Use:   "c6e1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run subcommand: clear/copy/len/remove")
	},
}

func init() {
	rootCmd.AddCommand(c6e1Cmd)

	c6e1IntSet.Add(1)
	c6e1IntSet.Add(2)
	c6e1IntSet.Add(3)
	c6e1IntSet.Add(4)

	// the issue is here, I should be creating a word list as a bit vector, which in this case is only one word:
	// ......0011111 (starting from one on the right across 64 values

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c6e1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c6e1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
