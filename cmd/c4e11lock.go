/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"the-go-programming-language/internal/composite-types/json"

	"github.com/spf13/cobra"
)

var c4e11LockIssueHeaderParams json.IssueHeaderParams
var c4e11LockIssuePathParams json.IssuePathParams
var c4e11LockIssueBodyParams json.IssueBodyParams

// lockCmd represents the create command
var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		json.RunLock(c4e11LockIssueHeaderParams, c4e11LockIssuePathParams, c4e11LockIssueBodyParams)
	},
}

func init() {
	C4e11Cmd.AddCommand(lockCmd)
	lockCmd.Flags().StringVarP(&c4e11LockIssueHeaderParams.Bearer, "bearer", "b", "", "Bearer token")
	lockCmd.Flags().StringVarP(&c4e11LockIssuePathParams.Repo, "repo", "r", "", "Repo the issue is in")
	lockCmd.Flags().StringVarP(&c4e11LockIssuePathParams.Owner, "owner", "o", "", "Owner of the repo the issue is in")
	lockCmd.Flags().StringVarP(&c4e11LockIssuePathParams.IssueNumber, "issue_number", "i", "", "Issue number of the issue")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
