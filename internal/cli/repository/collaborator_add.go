package repository

import (
	"fmt"
	"git-cli-client/internal/api"
	"git-cli-client/internal/utils"
	"github.com/spf13/cobra"
	"os"
)

var AddCollaboratorCmd = &cobra.Command{
	Use:   "collaborator-add",
	Short: "Add a collaborator to the repository",
	Run: func(cmd *cobra.Command, args []string) {
		owner, _ := cmd.Flags().GetString("owner")
		repo, _ := cmd.Flags().GetString("repo")
		user, _ := cmd.Flags().GetString("user")
		permission, _ := cmd.Flags().GetString("permission")

		addRepoMember(api.CollaboratorApi{
			Owner:      owner,
			Repo:       repo,
			User:       user,
			Permission: permission,
		})
	},
}

func init() {
	AddCollaboratorCmd.Flags().StringP("owner", "o", "", "Owner")
	AddCollaboratorCmd.Flags().StringP("repo", "r", "", "Repository")
	AddCollaboratorCmd.Flags().StringP("user", "u", "", "User")
	AddCollaboratorCmd.Flags().StringP("permission", "p", "", "Permission")

	if err := AddCollaboratorCmd.MarkFlagRequired("owner"); err != nil {
		return
	}
	if err := AddCollaboratorCmd.MarkFlagRequired("repo"); err != nil {
		return
	}
	if err := AddCollaboratorCmd.MarkFlagRequired("user"); err != nil {
		return
	}
}

func addRepoMember(ca api.CollaboratorApi) {
	client := api.NewClient(utils.GetToken())
	err := client.AddRepoMember(ca)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully collaborator added")
}
