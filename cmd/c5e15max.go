package cmd

import (
	"github.com/spf13/cobra"
	"the-go-programming-language/internal/functions/variadic"
)

var c5e15MaxTerms []int

// maxCmd represents the max command
var maxCmd = &cobra.Command{
	Use:   "max",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		variadic.Max(c5e15MaxTerms...)
	},
}

func init() {
	c5e15Cmd.AddCommand(maxCmd)
	maxCmd.Flags().IntSliceVarP(&c5e15MaxTerms, "terms", "t", []int{1, 2, 3, 4, 5}, "An comma separated string of integer search terms")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
