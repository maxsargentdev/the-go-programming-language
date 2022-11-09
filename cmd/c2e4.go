/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"the-go-programming-language/internal/program-structure/popcount"

	"github.com/spf13/cobra"
)

// c2e4Cmd represents the c2e4 command
var c2e4Cmd = &cobra.Command{
	Use:   "c2e4",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		popcount.LoopPopCount(1)
		popcount.BitMaskPopCount(1)
	},
}

func init() {
	rootCmd.AddCommand(c2e4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c2e4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c2e4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
