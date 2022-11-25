/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/composite-types/arrays"

	"github.com/spf13/cobra"
)

var c4e2inputString string
var c4e2inputHashFunc string

const (
	SHA256 string = "SHA256"
	SHA384 string = "SHA312"
	SHA512 string = "SHA512"
)

// c4e2Cmd represents the c4e2 command
var c4e2Cmd = &cobra.Command{
	Use:   "c4e2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch c4e2inputHashFunc {
		case SHA256:
			arrays.SHAOutput(c4e2inputString, SHA256)
		case SHA384:
			arrays.SHAOutput(c4e2inputString, SHA384)
		case SHA512:
			arrays.SHAOutput(c4e2inputString, SHA512)
		default:
			arrays.SHAOutput(c4e2inputString, SHA256)
		}
	},
}

func init() {
	rootCmd.AddCommand(c4e2Cmd)
	c4e2Cmd.Flags().StringVarP(&c4e2inputString, "input", "i", "hashme", "Input string for hash operation")
	c4e2Cmd.Flags().StringVarP(&c4e2inputHashFunc, "hashfunc", "f", "SHA256", "SHA256, SHA384 or SHA512")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c4e2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c4e2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
