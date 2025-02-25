package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [OPTIONS] [SERVICE...]",
	Short: "Creates containers for a service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().Bool("build", false, "Build images before starting containers.")
	createCmd.Flags().Bool("force-recreate", false, "Recreate containers even if their configuration and image haven't changed.")
	createCmd.Flags().Bool("no-build", false, "Don't build an image, even if it's missing.")
	createCmd.Flags().Bool("no-recreate", false, "If containers already exist, don't recreate them. Incompatible with --force-recreate.")
	createCmd.Flags().String("pull", "missing", "Pull image before running (\"always\"|\"missing\"|\"never\")")
	createCmd.Flags().Bool("remove-orphans", false, "Remove containers for services not defined in the Compose file.")
	createCmd.Flags().StringArray("scale", []string{}, "Scale SERVICE to NUM instances. Overrides the `scale` setting in the Compose file if present.")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"pull": carapace.ActionValues("always", "missing", "never").StyleF(style.ForKeyword),
	})

	carapace.Gen(createCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionServices(createCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
