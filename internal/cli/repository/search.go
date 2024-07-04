package repository

import (
	"fmt"
	"git-cli-client/internal/api"
	"git-cli-client/internal/utils"
	"github.com/spf13/cobra"
	"os"
)

var RepoSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "List all repositories",
	Run: func(cmd *cobra.Command, args []string) {
		SearchRepositories()
	},
}

func SearchRepositories() {
	client := api.NewClient(utils.GetToken())
	repos, err := client.GetRepositories()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, repo := range repos {
		fmt.Printf("ID: %d, Name: %s\n", repo.ID, repo.Name)
	}
}
