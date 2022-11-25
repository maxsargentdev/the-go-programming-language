/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/old/3-basic-data-types/anagram"

	"github.com/spf13/cobra"
)

var c3e12inputA string
var c3e12inputB string

// c3e12Cmd represents the c3e12 command
var c3e12Cmd = &cobra.Command{
	Use:   "c3e12",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		anagram.Anagram(c3e12inputA, c3e12inputB)
	},
}

func init() {
	rootCmd.AddCommand(c3e12Cmd)
	c3e12Cmd.Flags().StringVarP(&c3e12inputA, "inputA", "A", "racecar", "Input string for first half of anagram comparison")
	c3e12Cmd.Flags().StringVarP(&c3e12inputB, "inputB", "B", "racecar", "Input string for second half of anagram comparison")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c3e12Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c3e12Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
