/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/composite-types/slices"

	"github.com/spf13/cobra"
)

var c4e4inputSlice []int
var c4e4inputInt int

// c4e4Cmd represents the c4e4 command
var c4e4Cmd = &cobra.Command{
	Use:   "c4e4",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		slices.RotateLeft(c4e4inputInt, c4e4inputSlice)
	},
}

func init() {
	rootCmd.AddCommand(c4e4Cmd)
	c4e4Cmd.Flags().IntVarP(&c4e4inputInt, "rotatepositions", "r", 2, "How many positions to roteate")
	c4e4Cmd.Flags().IntSliceVarP(&c4e4inputSlice, "inputSLice", "i", []int{1, 2, 3, 4, 5}, "Input string slice for dedupe operation")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c4e4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c4e4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
