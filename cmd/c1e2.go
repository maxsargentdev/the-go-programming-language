package cmd

import (
	"github.com/spf13/cobra"
	"the-go-programming-language/chapters/1-tutorial/c1e2"
)

// c1e2Cmd represents the c1e2 command
var c1e2Cmd = &cobra.Command{
	Use:   "c1e2",
	Short: "Chapter 1 Exercise 2 - Echo with index and arguments",
	Long: `Chapter 1 Exercise 2 - Echo with index and arguments"

Description:
This command provides the similar functionality as the shell utility echo.
This command echos the index of the argument as well as the argument, on new lines.

Example:
./the-go-programming-language c1e2 <arguments to echo>
./the-go-programming-language c1e2 hello world

Notes:
Because we have completed this book as a CLI app, both the binary being executed & the subcommand c1e1 are echoed.
You can adjust this behaviour by manipulating the start index of the for loop.
`,
	Run: func(cmd *cobra.Command, args []string) {
		c1e2.EchoWithIndexAndArguments()
	},
}

func init() {
	rootCmd.AddCommand(c1e2Cmd)
}
