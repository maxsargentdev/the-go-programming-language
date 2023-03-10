package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"the-go-programming-language/internal/methods/intset"
)

// copyCmd copies the intset
var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		intSet := intset.IntSet{Words: c6e1Int64Set}
		fmt.Printf("%v\n", intSet)
		copiedIntSet := *intSet.Copy()
		fmt.Printf("%v\n", copiedIntSet)
	},
}

func init() {
	c6e1Cmd.AddCommand(copyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
