/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/functions/recursion"

	"github.com/spf13/cobra"
)

// c5e2Cmd represents the c5e2 command
var c5e2Cmd = &cobra.Command{
	Use:   "c5e2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		recursion.GenerateElementMap()
	},
}

func init() {
	rootCmd.AddCommand(c5e2Cmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c5e2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c5e2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
