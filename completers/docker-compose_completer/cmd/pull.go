package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull service images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().Bool("ignore-pull-failures", false, "Pull what it can and ignores images with pull failures.")
	pullCmd.Flags().Bool("include-deps", false, "Also pull services declared as dependencies")
	pullCmd.Flags().Bool("no-parallel", false, "Disable parallel pulling.")
	pullCmd.Flags().Bool("parallel", false, "Deprecated, pull multiple images in parallel (enabled by default).")
	pullCmd.Flags().BoolP("quiet", "q", false, "Pull without printing progress information")
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).PositionalAnyCompletion(
		action.ActionServices(pullCmd),
	)
}
