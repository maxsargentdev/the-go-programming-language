/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/interfaces/tempflag"

	"github.com/spf13/cobra"
)

// c7e7Cmd represents the c7e7 command
var c7e7Cmd = &cobra.Command{
	Use:   "c7e7",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tempflag.CobraOutput()
	},
}

func init() {
	rootCmd.AddCommand(c7e7Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c7e7Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c7e7Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
