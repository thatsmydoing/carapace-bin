package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_historyCmd = &cobra.Command{
	Use:   "history [OPTIONS] IMAGE",
	Short: "Show the history of an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_historyCmd).Standalone()
	image_historyCmd.Flags().String("format", "", "Format output using a custom template:")
	image_historyCmd.Flags().BoolP("human", "H", true, "Print sizes and dates in human readable format")
	image_historyCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	image_historyCmd.Flags().BoolP("quiet", "q", false, "Only show image IDs")
	imageCmd.AddCommand(image_historyCmd)

	rootAlias(image_historyCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(
			docker.ActionRepositoryTags(),
		)
	})
}
