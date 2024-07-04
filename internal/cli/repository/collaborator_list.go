package repository

import (
	"fmt"
	"git-cli-client/internal/api"
	"git-cli-client/internal/utils"
	"github.com/spf13/cobra"
	"os"
)

var CollaboratorListCmd = &cobra.Command{
	Use:   "collaborator-list",
	Short: "List all collaborators",
	Run: func(cmd *cobra.Command, args []string) {
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")

		repoMembers(api.CollaboratorApi{
			Owner: owner,
			Repo:  repo,
		})
	},
}

func init() {
	CollaboratorListCmd.Flags().StringP("owner", "o", "", "Owner")
	CollaboratorListCmd.Flags().StringP("repo", "r", "", "Repository")

	if err := CollaboratorListCmd.MarkFlagRequired("owner"); err != nil {
		return
	}
	if err := CollaboratorListCmd.MarkFlagRequired("repo"); err != nil {
		return
	}
}

func repoMembers(ca api.CollaboratorApi) {
	client := api.NewClient(utils.GetToken())
	users, err := client.GetRepoMembers(ca)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Username)
	}
}
