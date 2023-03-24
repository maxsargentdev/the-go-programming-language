/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"the-go-programming-language/internal/interfaces/counter"

	"github.com/spf13/cobra"
)

// c7e1Cmd represents the c7e1 command
var c7e1Cmd = &cobra.Command{
	Use:   "c7e1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var wc counter.WordCounter
		var lc counter.LineCounter

		var m string = `I am a multi line
			string
			count me
			in`

		// both types satisfy the io.Writer interface and thus we can use Fprintf
		fmt.Fprintf(&wc, "%s", m)
		fmt.Println(wc)

		fmt.Fprintf(&lc, "%s", m)
		fmt.Println(lc)
	},
}

func init() {
	rootCmd.AddCommand(c7e1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c7e1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c7e1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
