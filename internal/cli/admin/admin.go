package admin

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin commands",
}

func init() {
	Cmd.AddCommand(UserListCmd)
	Cmd.AddCommand(CreateUserCmd)
	Cmd.AddCommand(ActiveUserCmd)
	Cmd.AddCommand(InactiveUserCmd)
}
