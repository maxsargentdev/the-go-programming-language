/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"the-go-programming-language/internal/composite-types/maps"

	"github.com/spf13/cobra"
)

var c4e9InputFile string

// c4e9Cmd represents the c4e9 command
var c4e9Cmd = &cobra.Command{
	Use:   "c4e9",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if c4e9InputFile == "" {
			fmt.Println("Please provide a file path")
			os.Exit(0)
		}
		maps.Wordfreq(c4e9InputFile)
	},
}

func init() {
	rootCmd.AddCommand(c4e9Cmd)
	c4e9Cmd.Flags().StringVarP(&c4e9InputFile, "file", "f", "", "Input file for wordfreq")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c4e9Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c4e9Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
