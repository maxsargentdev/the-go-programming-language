/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"the-go-programming-language/internal/functions/panicandrecover"

	"github.com/spf13/cobra"
)

// c5e19Cmd represents the c5e19 command
var c5e19Cmd = &cobra.Command{
	Use:   "c5e19",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		result := panicandrecover.PanicAndRecover()
		fmt.Printf(result)
	},
}

func init() {
	rootCmd.AddCommand(c5e19Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c5e19Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c5e19Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
