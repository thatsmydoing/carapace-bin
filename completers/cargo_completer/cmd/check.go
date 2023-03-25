package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{"c"},
	Short:   "Check a local package and all of its dependencies for errors",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("check"),
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().Bool("all", false, "Alias for --workspace (deprecated)")
	checkCmd.Flags().Bool("all-features", false, "Activate all available features")
	checkCmd.Flags().Bool("all-targets", false, "Check all targets")
	checkCmd.Flags().StringSlice("bench", []string{}, "Check only the specified bench target")
	checkCmd.Flags().Bool("benches", false, "Check all benches")
	checkCmd.Flags().StringSlice("bin", []string{}, "Check only the specified binary")
	checkCmd.Flags().Bool("bins", false, "Check all binaries")
	checkCmd.Flags().StringSlice("example", []string{}, "Check only the specified example")
	checkCmd.Flags().Bool("examples", false, "Check all examples")
	checkCmd.Flags().StringSlice("exclude", []string{}, "Exclude packages from the check")
	checkCmd.Flags().StringSliceP("features", "F", []string{}, "Space or comma separated list of features to activate")
	checkCmd.Flags().Bool("future-incompat-report", false, "Outputs a future incompatibility report at the end of the build")
	checkCmd.Flags().BoolP("help", "h", false, "Print help")
	checkCmd.Flags().Bool("ignore-rust-version", false, "Ignore `rust-version` specification in packages")
	checkCmd.Flags().StringP("jobs", "j", "", "Number of parallel jobs, defaults to # of CPUs")
	checkCmd.Flags().Bool("keep-going", false, "Do not abort the build as soon as there is an error (unstable)")
	checkCmd.Flags().Bool("lib", false, "Check only this package's library")
	checkCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	checkCmd.Flags().StringSlice("message-format", []string{}, "Error format")
	checkCmd.Flags().Bool("no-default-features", false, "Do not activate the `default` feature")
	checkCmd.Flags().StringSliceP("package", "p", []string{}, "Package(s) to check")
	checkCmd.Flags().String("profile", "", "Check artifacts with the specified profile")
	checkCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	checkCmd.Flags().BoolP("release", "r", false, "Check artifacts in release mode, with optimizations")
	checkCmd.Flags().StringSlice("target", []string{}, "Check for the target triple")
	checkCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	checkCmd.Flags().StringSlice("test", []string{}, "Check only the specified test target")
	checkCmd.Flags().Bool("tests", false, "Check all tests")
	checkCmd.Flags().String("timings", "", "Timing output formats (unstable) (comma separated): html, json")
	checkCmd.Flags().Bool("unit-graph", false, "Output build graph in JSON (unstable)")
	checkCmd.Flags().Bool("workspace", false, "Check all packages in the workspace")
	checkCmd.Flag("timings").NoOptDefVal = " "
	rootCmd.AddCommand(checkCmd)

	carapace.Gen(checkCmd).FlagCompletion(carapace.ActionMap{
		"bench":          action.ActionTargets(checkCmd, action.TargetOpts{Bench: true}),
		"bin":            action.ActionTargets(checkCmd, action.TargetOpts{Bin: true}),
		"example":        action.ActionTargets(checkCmd, action.TargetOpts{Example: true}),
		"exclude":        action.ActionWorkspaceMembers(checkCmd),
		"features":       action.ActionFeatures(checkCmd).UniqueList(","),
		"manifest-path":  carapace.ActionFiles(),
		"message-format": action.ActionMessageFormats(),
		"package":        action.ActionDependencies(buildCmd, true),
		"profile":        action.ActionProfiles(checkCmd),
		"target-dir":     carapace.ActionDirectories(),
		"test":           action.ActionTargets(checkCmd, action.TargetOpts{Test: true}),
	})
}
