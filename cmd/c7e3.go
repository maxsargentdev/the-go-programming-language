/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/interfaces/counter"

	"github.com/spf13/cobra"
)

// c7e3Cmd represents the c7e3 command
var c7e3Cmd = &cobra.Command{
	Use:   "c7e3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tree := counter.Tree{}
		tree.Value = 1
		tree.Right = &counter.Tree{Value: 2}
		tree.Left = &counter.Tree{Value: 3}
		tree.Left.Left = &counter.Tree{Value: 10}
		tree.String()

	},
}

func init() {
	rootCmd.AddCommand(c7e3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c7e3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c7e3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
