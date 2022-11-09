/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"the-go-programming-language/internal/program-structure/conv"

	"github.com/spf13/cobra"
)

var c2e2measures []string

// c2e2Cmd represents the c2e2 command
var c2e2Cmd = &cobra.Command{
	Use:   "c2e2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		conv.Convert(c2e2measures)
	},
}

func init() {
	rootCmd.AddCommand(c2e2Cmd)
	c2e2Cmd.Flags().StringSliceVarP(&c2e2measures, "measures", "m", []string{"100°C", "100°F", "100Ft", "100M", "100lbs", "100kgs"}, "An comma separated string of measures") // need a much longer list to go here

	// would we want to turn off the above defaults so that we can also use std in

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c2e2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c2e2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
