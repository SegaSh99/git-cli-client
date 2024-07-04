package admin

import (
	"fmt"
	"git-cli-client/internal/api"
	"git-cli-client/internal/utils"
	"git-cli-client/pkg/models"
	"github.com/spf13/cobra"
	"os"
)

var InactiveUserCmd = &cobra.Command{
	Use:   "user-inactive",
	Short: "Inactivate the user",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")

		inactiveUser(models.UpdateUser{
			Username: username,
			Active:   false,
		})
	},
}

func init() {
	InactiveUserCmd.Flags().StringP("username", "u", "", "Username")

	if err := InactiveUserCmd.MarkFlagRequired("username"); err != nil {
		return
	}
}

func inactiveUser(uu models.UpdateUser) {
	client := api.NewClient(utils.GetToken())
	if err := client.UpdateUser(uu); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("User inactivated successfully")
}
