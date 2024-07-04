package auth

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "auth",
	Short: "Authorization/registration commands",
}

func init() {
	Cmd.AddCommand(SignInCmd)
}
