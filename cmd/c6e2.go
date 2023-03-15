/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"the-go-programming-language/internal/methods/intset"

	"github.com/spf13/cobra"
)

var c6e2IntSlice []int
var c6e2IntSet intset.IntSet

// c6e2Cmd represents the c6e2 command
var c6e2Cmd = &cobra.Command{
	Use:   "c6e2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(c6e2IntSet.String())
		c6e2IntSet.AddAll([]int{1, 2, 3, 4, 5}...)
		fmt.Println(c6e2IntSet.String())
	},
}

func init() {
	rootCmd.AddCommand(c6e2Cmd)
	c6e2Cmd.Flags().IntSliceVarP(&c6e2IntSlice, "intset", "i", []int{1, 1, 2, 3, 3}, "An comma separated string of integer elements")

	for _, v := range c6e2IntSlice {
		c6e2IntSet.Add(v)
	}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c6e2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c6e2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
