/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"the-go-programming-language/internal/methods/intset"

	"github.com/spf13/cobra"
)

var c6e3IntSet intset.IntSet

// c6e3Cmd represents the c6e3 command
var c6e3Cmd = &cobra.Command{
	Use:   "c6e3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("c6e3 called")
	},
}

func init() {
	rootCmd.AddCommand(c6e3Cmd)

	c6e3IntSet.Add(5)
	c6e3IntSet.Add(6)
	c6e3IntSet.Add(7)
	c6e3IntSet.Add(8)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c6e3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c6e3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
