/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"the-go-programming-language/internal/tutorial"

	"github.com/spf13/cobra"
)

// c1e4Cmd represents the c1e4 command
var c1e4Cmd = &cobra.Command{
	Use:   "c1e4",
	Short: "Chapter 1 Exercise 4 - Modified dup2, print names of all files with duplicate lines",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tutorial.Dup2()
	},
}

func init() {
	rootCmd.AddCommand(c1e4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c1e4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c1e4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
