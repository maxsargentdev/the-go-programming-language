package cmd

import (
	"github.com/spf13/cobra"
	"the-go-programming-language/internal/methods/intset"
)

// clearCmd clears the intset
var intersectCmd = &cobra.Command{
	Use:   "intersect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		testIntSet := intset.IntSet{}
		testIntSet.Add(5)
		c6e3IntSet.IntersectWith(&testIntSet)
		println(c6e3IntSet.String())
	},
}

func init() {
	c6e3Cmd.AddCommand(intersectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
