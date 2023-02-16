/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"the-go-programming-language/internal/functions/functionvalues"
)

var c5e8URL string
var c5e8ID string

// c5e8Cmd represents the c5e8 command
var c5e8Cmd = &cobra.Command{
	Use:   "c5e8",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		functionvalues.RunGetElementByID(c5e8URL, c5e8ID)
	},
}

func init() {
	rootCmd.AddCommand(c5e8Cmd)
	c5e8Cmd.Flags().StringVarP(&c5e8URL, "url", "u", "https://xkcd.com", "the url to use for running pretty print")
	c5e8Cmd.Flags().StringVarP(&c5e8ID, "id", "i", "topContainer", "the id to lookup on the page")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c5e8Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c5e8Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
