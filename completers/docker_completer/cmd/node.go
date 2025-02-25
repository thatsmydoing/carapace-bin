package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var nodeCmd = &cobra.Command{
	Use:     "node",
	Short:   "Manage Swarm nodes",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nodeCmd).Standalone()
	rootCmd.AddCommand(nodeCmd)
}
