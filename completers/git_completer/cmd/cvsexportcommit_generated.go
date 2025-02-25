package cmd

import (
	"github.com/spf13/cobra"
)

var cvsexportcommitCmd = &cobra.Command{
	Use:     "cvsexportcommit",
	Short:   "Export a single commit to a CVS checkout",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {

	rootCmd.AddCommand(cvsexportcommitCmd)
}
