/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/composite-types/json"

	"github.com/spf13/cobra"
)

var c4e14Project string
var c4e14Repo string

// c4e14Cmd represents the c4e14 command
var c4e14Cmd = &cobra.Command{
	Use:   "c4e14",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		json.Serve(c4e14Project, c4e14Repo)
	},
}

func init() {
	rootCmd.AddCommand(c4e14Cmd)
	c4e14Cmd.Flags().StringVarP(&c4e14Project, "project", "p", "golang", "project that owns this repo")
	c4e14Cmd.Flags().StringVarP(&c4e14Repo, "repo", "r", "go", "repo name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// c4e14Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// c4e14Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
