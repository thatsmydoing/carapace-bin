package cmd

import (
	"github.com/spf13/cobra"
)

var index_packCmd = &cobra.Command{
	Use:     "index-pack",
	Short:   "Build pack index file for an existing packed archive",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {

	rootCmd.AddCommand(index_packCmd)
}
