/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/functions/functionvalues"

	"github.com/spf13/cobra"
)

var c5e7URL string

// c5e7Cmd represents the c5e7 command
var c5e7Cmd = &cobra.Command{
	Use:   "c5e7",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		functionvalues.RunHTMLPrettyPrint(c5e7URL)
	},
}

func init() {
	rootCmd.AddCommand(c5e7Cmd)
	c5e7Cmd.Flags().StringVarP(&c5e7URL, "url", "u", "https://xkcd.com", "the url to use for running pretty print")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c5e7Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c5e7Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
