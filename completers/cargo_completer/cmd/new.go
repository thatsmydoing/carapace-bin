package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "Create a new cargo package at <path>",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("new"),
}

func init() {
	carapace.Gen(newCmd).Standalone()

	newCmd.Flags().Bool("bin", false, "Use a binary (application) template [default]")
	newCmd.Flags().String("edition", "", "Edition to set for the crate generated")
	newCmd.Flags().BoolP("help", "h", false, "Print help")
	newCmd.Flags().Bool("lib", false, "Use a library template")
	newCmd.Flags().String("name", "", "Set the resulting package name, defaults to the directory name")
	newCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	newCmd.Flags().String("registry", "", "Registry to use")
	newCmd.Flags().String("vcs", "", "Initialize a new repository for the given version control system (git, hg, pijul, or fossil) or do not initialize any version control at all (none), overriding a global configuration.")
	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).FlagCompletion(carapace.ActionMap{
		"edition":  carapace.ActionValues("2015", "2018"),
		"registry": action.ActionRegistries(),
		"vcs":      carapace.ActionValues("git", "hg", "pijul", "vcs", "none"),
	})

	carapace.Gen(newCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
