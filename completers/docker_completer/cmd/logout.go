package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout [SERVER]",
	Short: "Log out from a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logoutCmd).Standalone()
	rootCmd.AddCommand(logoutCmd)
}
