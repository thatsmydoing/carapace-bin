package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker-compose_completer/cmd/action"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:     "config [OPTIONS] [SERVICE...]",
	Short:   "Parse, resolve and render compose file in canonical format",
	Aliases: []string{"convert"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).Standalone()

	configCmd.Flags().String("format", "yaml", "Format the output. Values: [yaml | json]")
	configCmd.Flags().String("hash", "", "Print the service config hash, one per line.")
	configCmd.Flags().Bool("images", false, "Print the image names, one per line.")
	configCmd.Flags().Bool("no-consistency", false, "Don't check model consistency - warning: may produce invalid Compose output")
	configCmd.Flags().Bool("no-interpolate", false, "Don't interpolate environment variables.")
	configCmd.Flags().Bool("no-normalize", false, "Don't normalize compose model.")
	configCmd.Flags().StringP("output", "o", "", "Save to file (default to stdout)")
	configCmd.Flags().Bool("profiles", false, "Print the profile names, one per line.")
	configCmd.Flags().BoolP("quiet", "q", false, "Only validate the configuration, don't print anything.")
	configCmd.Flags().Bool("resolve-image-digests", false, "Pin image tags to digests.")
	configCmd.Flags().Bool("services", false, "Print the service names, one per line.")
	configCmd.Flags().Bool("volumes", false, "Print the volume names, one per line.")
	rootCmd.AddCommand(configCmd)

	carapace.Gen(configCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("yaml", "json"),
		"hash":   action.ActionServices(configCmd),
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(configCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionServices(configCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
