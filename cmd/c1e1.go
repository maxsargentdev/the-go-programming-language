package cmd

import (
	"github.com/spf13/cobra"
	"the-go-programming-language/chapters/1-tutorial/c1e1"
)

// c1e1Cmd represents the c1e1 command
var c1e1Cmd = &cobra.Command{
	Use:   "c1e1",
	Short: "Chapter 1 Exercise 1 - Echo with command name",
	Long: `Chapter 1 Exercise 1 - Echo with command name

This command provides the similar functionality as the shell utility echo,
minus the fact that this command also echos the name of the command.

Example:
./the-go-programming-language c1e1 <arguments to echo>
./the-go-programming-language c1e1 hello world
`,
	Run: func(cmd *cobra.Command, args []string) {
		c1e1.EchoWithCommandName()
	},
}

func init() {
	rootCmd.AddCommand(c1e1Cmd)
}
