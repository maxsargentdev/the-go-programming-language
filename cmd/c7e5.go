/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"the-go-programming-language/internal/interfaces/parser"

	"github.com/spf13/cobra"
)

// c7e5Cmd represents the c7e5 command
var c7e5Cmd = &cobra.Command{
	Use:   "c7e5",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := "abcdefghijklmnopqrtsuvwxyz"
		b := &bytes.Buffer{}
		r := parser.LimitReader(strings.NewReader(s), 4)
		n, _ := b.ReadFrom(r)
		fmt.Printf("%d bytes read: %s", n, b.String())
	},
}

func init() {
	rootCmd.AddCommand(c7e5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c7e5Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c7e5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
