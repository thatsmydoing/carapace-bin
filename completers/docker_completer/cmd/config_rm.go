package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more configs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_rmCmd).Standalone()

	configCmd.AddCommand(config_rmCmd)

	carapace.Gen(config_rmCmd).PositionalAnyCompletion(action.ActionConfigs())
}
