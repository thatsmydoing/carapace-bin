package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/supervisor"
	"github.com/spf13/cobra"
)

var signalCmd = &cobra.Command{
	Use:   "signal",
	Short: "Signal process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(signalCmd).Standalone()

	rootCmd.AddCommand(signalCmd)

	carapace.Gen(signalCmd).PositionalCompletion(
		supervisor.ActionSignals(),
	)
}
