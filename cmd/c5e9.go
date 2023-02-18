/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/functions/functionvalues"

	"github.com/spf13/cobra"
)

var c5e9String string

// c5e9Cmd represents the c5e9 command
var c5e9Cmd = &cobra.Command{
	Use:   "c5e9",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		functionvalues.RunExpand(c5e9String)
	},
}

func init() {
	rootCmd.AddCommand(c5e9Cmd)
	c5e9Cmd.Flags().StringVarP(&c5e9String, "string", "s", "I am a string, my favourite $foo is foo.", "the string to use for interpolation testing")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c5e9Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c5e9Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
