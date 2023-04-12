/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"the-go-programming-language/internal/interfaces/tempflag"
)

// c7e6Cmd represents the c7e6 command
var c7e6Cmd = &cobra.Command{
	Use:   "c7e6",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(c7e6Cmd)

	// should move this into tempflag.CelsiusFlag fo complete example properly
	c7e6Cmd.Flags().AddFlag(&pflag.Flag{
		Name:      "temp",
		Shorthand: "t",
		Usage:     "convert any temperature to its celsius value",
		Value:     tempflag.CelsiusFlagType{},
		DefValue:  "20C",
	})

	c7e6Cmd.Flags().Parse([]string{"temp"})
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c7e6Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c7e6Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
