package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [OPTIONS] SERVICE [COMMAND] [ARGS...]",
	Short: "Run a one-off command on a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("build", false, "Build image before starting container.")
	runCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")
	runCmd.Flags().String("entrypoint", "", "Override the entrypoint of the image")
	runCmd.Flags().StringArrayP("env", "e", []string{}, "Set environment variables")
	runCmd.Flags().BoolP("interactive", "i", true, "Keep STDIN open even if not attached.")
	runCmd.Flags().StringArrayP("label", "l", []string{}, "Add or override a label")
	runCmd.Flags().String("name", "", "Assign a name to the container")
	runCmd.Flags().BoolP("no-TTY", "T", false, "Disable pseudo-TTY allocation (default: auto-detected).")
	runCmd.Flags().Bool("no-deps", false, "Don't start linked services.")
	runCmd.Flags().StringArrayP("publish", "p", []string{}, "Publish a container's port(s) to the host.")
	runCmd.Flags().Bool("quiet-pull", false, "Pull without printing progress information.")
	runCmd.Flags().Bool("remove-orphans", false, "Remove containers for services not defined in the Compose file.")
	runCmd.Flags().Bool("rm", false, "Automatically remove the container when it exits")
	runCmd.Flags().Bool("service-ports", false, "Run command with the service's ports enabled and mapped to the host.")
	runCmd.Flags().BoolP("tty", "t", true, "Allocate a pseudo-TTY.")
	runCmd.Flags().Bool("use-aliases", false, "Use the service's network useAliases in the network(s) the container connects to.")
	runCmd.Flags().StringP("user", "u", "", "Run as specified username or uid")
	runCmd.Flags().StringArrayP("volume", "v", []string{}, "Bind mount a volume.")
	runCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	rootCmd.AddCommand(runCmd)

	// TODO flag completion
	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"volume": action.ActionVolumes(runCmd),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		action.ActionServices(runCmd),
	)
}
