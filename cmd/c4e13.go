/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/composite-types/json"

	"github.com/spf13/cobra"
)

var c4e13SearchTerm string

// c4e13Cmd represents the c4e13 command
var c4e13Cmd = &cobra.Command{
	Use:   "c4e13",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		json.RunOMDBFetch(c4e13SearchTerm)
	},
}

func init() {
	rootCmd.AddCommand(c4e13Cmd)
	c4e13Cmd.Flags().StringVarP(&c4e13SearchTerm, "searchTerm", "s", "Star Wars: Episode IV - A New Hope", "Search term for OMDB")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c4e13Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c4e13Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
