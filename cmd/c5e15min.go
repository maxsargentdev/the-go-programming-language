package cmd

import (
	"github.com/spf13/cobra"
	"the-go-programming-language/internal/functions/variadic"
)

var c5e15MinTerms []int

// minCmd represents the max command
var minCmd = &cobra.Command{
	Use:   "min",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		variadic.Min(c5e15MinTerms...)
	},
}

func init() {
	c5e15Cmd.AddCommand(minCmd)
	minCmd.Flags().IntSliceVarP(&c5e15MinTerms, "terms", "t", []int{1, 2, 3, 4, 5}, "An comma separated string of integer search terms")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
