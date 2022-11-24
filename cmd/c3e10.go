/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"the-go-programming-language/internal/basic-data-types/string/comma"

	"github.com/spf13/cobra"
)

var c3e10input string

// c3e10Cmd represents the c3e10 command
var c3e10Cmd = &cobra.Command{
	Use:   "c3e10",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		comma.Comma(c3e10input)
	},
}

func init() {
	rootCmd.AddCommand(c3e10Cmd)
	c3e10Cmd.Flags().StringVarP(&c3e10input, "input", "i", "123456789", "Input number (string representation) to the comma function") // need a much longer list to go here

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c3e10Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c3e10Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
