package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var findDupesCmd = &cobra.Command{
	Use:   "find-dupes",
	Short: "Find duplication in the package tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findDupesCmd).Standalone()
	findDupesCmd.Flags().Bool("bin-links", false, "create symlinks")
	findDupesCmd.Flags().Bool("fund", false, "Display funding message")
	findDupesCmd.Flags().Bool("global-style", false, "Use global layout")
	findDupesCmd.Flags().Bool("ignore-scripts", false, "Disable scripts")
	findDupesCmd.Flags().Bool("legacy-bundling", false, "Use legacy bundling")
	findDupesCmd.Flags().StringArray("omit", []string{}, "omit dependency type")
	findDupesCmd.Flags().Bool("package-lock", false, "When false ignore package.lock")
	findDupesCmd.Flags().Bool("strict-peer-deps", false, "Fail and abort for any conflicting `peerDependencies`")
	findDupesCmd.Flags().StringP("workspace", "w", "", "Enable running a command in the context of the given workspace")
	findDupesCmd.Flags().Bool("workspaces", false, "Enable running a command in the context fo all workspaces")

	rootCmd.AddCommand(findDupesCmd)

	carapace.Gen(findDupesCmd).FlagCompletion(carapace.ActionMap{
		"omit": carapace.ActionValues("dev", "optional", "peer"),
	})
}
