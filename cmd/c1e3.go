/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"the-go-programming-language/chapters/1-tutorial/c1e3"
)

// c1e3Cmd represents the c1e3 command
var c1e3Cmd = &cobra.Command{
	Use:   "c1e3",
	Short: "Chapter 1 Exercise 3 - Comparing the performance of two echo functions",
	Long: `Chapter 1 Exercise 3 - Comparing the performance of two echo functions"

Description:
This command compares the speed of two different implementations of an echo.
The first recreates the string each iteration causing some wasted memory usage.
The second uses the strings.Join which concatenates the existing strings.

Example:
./the-go-programming-language c1e3 <arguments to echo>
./the-go-programming-language c1e3 hello world

Notes:
Because we have completed this book as a CLI app, both the binary being executed & the subcommand c1e1 were echoed.
We adjusted this behaviour by manipulating the os.Args slice and the loop index.

To-Do:
Upgrade this command to take a flag that specifies whether or not to take the binary being executed as an item.
`,
	Run: func(cmd *cobra.Command, args []string) {
		c1e3.TimedEchoTest()
	},
}

func init() {
	rootCmd.AddCommand(c1e3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c1e3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c1e3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
