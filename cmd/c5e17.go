/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/functions/variadic"

	"github.com/spf13/cobra"
)

var c5e17URL string
var c5e17Terms []string

// c5e17Cmd represents the c5e17 command
var c5e17Cmd = &cobra.Command{
	Use:   "c5e17",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		variadic.RunGetElementsByTagName(c5e17URL, c5e17Terms...)
	},
}

func init() {
	rootCmd.AddCommand(c5e17Cmd)
	c5e17Cmd.Flags().StringVarP(&c5e17URL, "url", "u", "https://xkcd.com", "the url to use")
	c5e17Cmd.Flags().StringSliceVarP(&c5e17Terms, "terms", "t", []string{"div"}, "the search terms")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c5e17Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c5e17Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
