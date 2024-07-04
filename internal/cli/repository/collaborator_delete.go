package repository

import (
	"fmt"
	"git-cli-client/internal/api"
	"git-cli-client/internal/utils"
	"github.com/spf13/cobra"
	"os"
)

var DeleteCollaboratorCmd = &cobra.Command{
	Use:   "collaborator-delete",
	Short: "Delete a collaborator to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		user, _ := cmd.Flags().GetString("user")

		deleteRepoMember(api.CollaboratorApi{
			Owner: owner,
			Repo:  repo,
			User:  user,
		})
	},
}

func init() {
	DeleteCollaboratorCmd.Flags().StringP("owner", "o", "", "Owner")
	DeleteCollaboratorCmd.Flags().StringP("repo", "r", "", "Repository")
	DeleteCollaboratorCmd.Flags().StringP("user", "u", "", "User")

	if err := DeleteCollaboratorCmd.MarkFlagRequired("owner"); err != nil {
		return
	}
	if err := DeleteCollaboratorCmd.MarkFlagRequired("repo"); err != nil {
		return
	}
	if err := DeleteCollaboratorCmd.MarkFlagRequired("user"); err != nil {
		return
	}
}

func deleteRepoMember(ca api.CollaboratorApi) {
	client := api.NewClient(utils.GetToken())
	err := client.DeleteRepoMember(ca)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully collaborator deleted")
}
