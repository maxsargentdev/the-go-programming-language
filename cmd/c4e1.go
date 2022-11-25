/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/composite-types/arrays"

	"github.com/spf13/cobra"
)

var c4e1inputA string
var c4e1inputB string

// c4e1Cmd represents the c4e1 command
var c4e1Cmd = &cobra.Command{
	Use:   "c4e1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		arrays.SHABitCompare(c4e1inputA, c4e1inputB)
	},
}

func init() {
	rootCmd.AddCommand(c4e1Cmd)
	c4e1Cmd.Flags().StringVarP(&c4e1inputA, "inputA", "A", "X", "Input string for first half of SHA checksum comparison")
	c4e1Cmd.Flags().StringVarP(&c4e1inputB, "inputB", "B", "X", "Input string for second half of SHA checksum comparison")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c4e1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c4e1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
