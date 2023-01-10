/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/composite-types/json"

	"github.com/spf13/cobra"
)

var c4e11UpdateIssueHeaderParams json.IssueHeaderParams
var c4e11UpdateIssuePathParams json.IssuePathParams
var c4e11UpdateIssueBodyParams json.IssueBodyParams

// updateCmd represents the create command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		json.RunUpdate(c4e11UpdateIssueHeaderParams, c4e11UpdateIssuePathParams, c4e11UpdateIssueBodyParams)
	},
}

func init() {
	C4e11Cmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&c4e11UpdateIssueHeaderParams.Bearer, "bearer", "b", "", "Bearer token")

	updateCmd.Flags().StringVarP(&c4e11UpdateIssuePathParams.Repo, "repo", "r", "", "Repo to update the issue from")
	updateCmd.Flags().StringVarP(&c4e11UpdateIssuePathParams.Owner, "owner", "o", "", "Owner of the repo")
	updateCmd.Flags().StringVarP(&c4e11UpdateIssuePathParams.IssueNumber, "issue_number", "i", "", "Issue number of the issue")

	updateCmd.Flags().StringVarP(&c4e11UpdateIssueBodyParams.Title, "title", "t", "", "Title of the issue")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
