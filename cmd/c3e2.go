/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"the-go-programming-language/internal/basic-data-types/floatingp"

	"github.com/spf13/cobra"
)

// c3e2Cmd represents the c3e2 command
var c3e2Cmd = &cobra.Command{
	Use:   "c3e2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		floatingp.Surface("eggbox", false)
	},
}

func init() {
	rootCmd.AddCommand(c3e2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c3e2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c3e2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
