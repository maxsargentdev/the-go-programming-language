/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"the-go-programming-language/internal/tutorial"

	"github.com/spf13/cobra"
)

// c1e7Cmd represents the c1e7 command
var c1e7Cmd = &cobra.Command{
	Use:   "c1e7",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("c1e7 called")
		tutorial.BasicFetch()
	},
}

func init() {
	rootCmd.AddCommand(c1e7Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c1e7Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c1e7Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
