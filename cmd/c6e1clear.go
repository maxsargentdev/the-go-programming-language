package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"the-go-programming-language/internal/methods/intset"
)

// clearCmd clears the intset
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		intSet := intset.IntSet{}
		intSet.Add(1)
		intSet.Add(44)
		intSet.Add(9)
		fmt.Println(intSet.String())
		intSet.Clear()
		fmt.Println(intSet.String())
	},
}

func init() {
	c6e1Cmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
