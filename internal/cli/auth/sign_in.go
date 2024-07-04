package auth

import (
	"fmt"
	"git-cli-client/internal/api"
	"github.com/spf13/cobra"
	"os"
)

var SignInCmd = &cobra.Command{
	Use:   "sign-in",
	Short: "User authorization",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		authenticate(username, password)
	},
}

func init() {
	SignInCmd.Flags().StringP("username", "u", "", "Username")
	SignInCmd.Flags().StringP("password", "p", "", "Password")

	if err := SignInCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := SignInCmd.MarkFlagRequired("password"); err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func authenticate(username, password string) {
	client := api.NewClient("")
	_, err := client.Login(username, password)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Successfully authorisation")
}
