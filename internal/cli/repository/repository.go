package repository

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "repository",
	Short: "Repository management",
}

func init() {
	Cmd.AddCommand(RepoSearchCmd)
	Cmd.AddCommand(CollaboratorListCmd)
	Cmd.AddCommand(AddCollaboratorCmd)
	Cmd.AddCommand(DeleteCollaboratorCmd)
}
