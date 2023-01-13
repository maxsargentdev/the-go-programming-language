/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/composite-types/json"

	"github.com/spf13/cobra"
)

// materializeCmd represents the materialize command
var materializeCmd = &cobra.Command{
	Use:   "materialize",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		json.RunXKCDIndexGen()
		json.RunXKCDMaterialize()
	},
}

func init() {
	c4e12Cmd.AddCommand(materializeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// materializeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// materializeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
