/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var c6e1IntSet []uint
var c6e1Int64Set []uint64

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
	c6e1Cmd.Flags().UintSliceVarP(&c6e1IntSet, "intset", "i", []uint{1, 2, 3, 4, 5}, "An comma separated string of integer elements")

	for _, v := range c6e1IntSet {
		c6e1Int64Set = append(c6e1Int64Set, uint64(v))
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c6e1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c6e1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
