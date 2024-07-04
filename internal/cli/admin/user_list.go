package admin

import (
	"fmt"
	"git-cli-client/internal/api"
	"git-cli-client/internal/utils"
	"os"

	"github.com/spf13/cobra"
)

var (
	userStatus  string
	UserListCmd = &cobra.Command{
		Use:   "user-list",
		Short: "List all users",
		Run: func(cmd *cobra.Command, args []string) {
			listUsers()
		},
	}
)

func init() {
	UserListCmd.Flags().StringVarP(&userStatus, "active", "a", "", "Filter users by status (active/inactive)")
}

func listUsers() {
	client := api.NewClient(utils.GetToken())
	users, err := client.GetUsers()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	for _, user := range users {
		if userStatus == "" || (userStatus == "true" && user.Active) || (userStatus == "false" && !user.Active) {
			fmt.Printf("ID: %d, Login: %s, Email: %s, Active: %v\n", user.ID, user.Login, user.Email, user.Active)
		}
	}
}
