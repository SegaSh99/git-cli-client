package admin

import (
	"fmt"
	"git-cli-client/internal/api"
	"git-cli-client/internal/utils"
	"git-cli-client/pkg/models"
	"github.com/spf13/cobra"
	"os"
)

var ActiveUserCmd = &cobra.Command{
	Use:   "user-active",
	Short: "Activate the user",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")

		activeUser(models.UpdateUser{
			Username: username,
			Active:   true,
		})
	},
}

func init() {
	ActiveUserCmd.Flags().StringP("username", "u", "", "Username")

	if err := ActiveUserCmd.MarkFlagRequired("username"); err != nil {
		return
	}
}

func activeUser(uu models.UpdateUser) {
	client := api.NewClient(utils.GetToken())
	if err := client.UpdateUser(uu); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("User activated successfully")
}
