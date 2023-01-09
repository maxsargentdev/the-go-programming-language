/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/composite-types/json"

	"github.com/spf13/cobra"
)

var c4e11CreateIssueHeaderParams json.IssueHeaderParams
var c4e11CreateIssuePathParams json.IssuePathParams
var c4e11CreateIssueBodyParams json.IssueBodyParams

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		json.RunCreate(c4e11CreateIssueHeaderParams, c4e11CreateIssuePathParams, c4e11CreateIssueBodyParams)
	},
}

func init() {
	C4e11Cmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&c4e11CreateIssueHeaderParams.Bearer, "bearer", "b", "", "Bearer token")

	createCmd.Flags().StringVarP(&c4e11CreateIssuePathParams.Repo, "repo", "r", "", "Repo to create the issue in")
	createCmd.Flags().StringVarP(&c4e11CreateIssuePathParams.Owner, "owner", "o", "", "Owner of the repo to create the issue in")

	createCmd.Flags().StringVarP(&c4e11CreateIssueBodyParams.Title, "title", "t", "", "Title of the issue")

	// Add body params here??

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
