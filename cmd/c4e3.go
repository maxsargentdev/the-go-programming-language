/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/composite-types/slices"

	"github.com/spf13/cobra"
)

// c4e3Cmd represents the c4e3 command
var c4e3Cmd = &cobra.Command{
	Use:   "c4e3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// slices.Reverse([]int{1, 2, 3, 4, 5})
		array := [5]int{1, 2, 3, 4, 5} // todo - this needs to be a arg, but its an array type so lets leave for now, pflags only uses slices
		slices.ReverseUsingArrayPointer(&array)

	},
}

func init() {
	rootCmd.AddCommand(c4e3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c4e3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c4e3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
