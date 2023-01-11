/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"the-go-programming-language/internal/composite-types/json"
)

var c4e11ReadIssueHeaderParams json.IssueHeaderParams
var c4e11ReadIssuePathParams json.IssuePathParams
var c4e11ReadIssueBodyParams json.IssueBodyParams

// createCmd represents the create command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		json.RunGitHubRead(c4e11ReadIssueHeaderParams, c4e11ReadIssuePathParams, c4e11ReadIssueBodyParams)
	},
}

func init() {
	C4e11Cmd.AddCommand(readCmd)
	readCmd.Flags().StringVarP(&c4e11ReadIssueHeaderParams.Bearer, "bearer", "b", "", "Bearer token")

	readCmd.Flags().StringVarP(&c4e11ReadIssuePathParams.Repo, "repo", "r", "", "Repo the issue is in")
	readCmd.Flags().StringVarP(&c4e11ReadIssuePathParams.Owner, "owner", "o", "", "Owner of the repo")
	readCmd.Flags().StringVarP(&c4e11ReadIssuePathParams.IssueNumber, "issue_number", "i", "", "Issue number of the issue")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
