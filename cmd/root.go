package cmd

import (
	"fmt"
	"git-cli-client/internal/cli/admin"
	"git-cli-client/internal/cli/auth"
	"git-cli-client/internal/cli/repository"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "CLI client for server administration",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(auth.Cmd)
	rootCmd.AddCommand(repository.Cmd)
	rootCmd.AddCommand(admin.Cmd)
}
