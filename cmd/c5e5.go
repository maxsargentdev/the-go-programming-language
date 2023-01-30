/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/functions/multireturn"

	"github.com/spf13/cobra"
)

var c5e5URL string

// c5e5Cmd represents the c5e5 command
var c5e5Cmd = &cobra.Command{
	Use:   "c5e5",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		multireturn.RunCountWordsAndImages(c5e5URL)
	},
}

func init() {
	rootCmd.AddCommand(c5e5Cmd)
	c5e5Cmd.Flags().StringVarP(&c5e5URL, "url", "u", "https://www.xkcd.com", "url to use for word and image count")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c5e5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c5e5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
