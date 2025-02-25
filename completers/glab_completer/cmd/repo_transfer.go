package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_transferCmd = &cobra.Command{
	Use:   "transfer [repo] [flags]",
	Short: "Transfer a repository to a new namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_transferCmd).Standalone()
	repo_transferCmd.Flags().StringP("target-namespace", "t", "", "The namespace where your project should be transferred to")
	repo_transferCmd.Flags().BoolP("yes", "y", false, "Danger: Skip confirmation prompt and force transfer operation. Transfer cannot be undone.")
	repoCmd.AddCommand(repo_transferCmd)

	// TODO target-namespace completion (user/group ?)

	carapace.Gen(repo_transferCmd).PositionalCompletion(
		action.ActionRepo(repo_transferCmd), // TODO verify this is correct
	)
}
