package admin

import (
	"fmt"
	"git-cli-client/internal/api"
	"git-cli-client/internal/utils"
	"git-cli-client/pkg/models"
	"github.com/spf13/cobra"
	"os"
)

var CreateUserCmd = &cobra.Command{
	Use:   "user-create",
	Short: "Create a new user",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")

		createUser(models.CreateUser{
			Username: username,
			Email:    email,
			Password: password,
		})
	},
}

func init() {
	CreateUserCmd.Flags().StringP("username", "u", "", "Username")
	CreateUserCmd.Flags().StringP("email", "e", "", "Email")
	CreateUserCmd.Flags().StringP("password", "p", "", "Password")

	if err := CreateUserCmd.MarkFlagRequired("username"); err != nil {
		return
	}
	if err := CreateUserCmd.MarkFlagRequired("email"); err != nil {
		return
	}
	if err := CreateUserCmd.MarkFlagRequired("password"); err != nil {
		return
	}
}

func createUser(cu models.CreateUser) {
	client := api.NewClient(utils.GetToken())
	if err := client.CreateUser(cu); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("User created successfully")
}
