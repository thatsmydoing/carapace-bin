package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "kubectl controls the Kubernetes cluster manager",
	Long:  "https://kubernetes.io/docs/reference/kubectl/overview/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "basic beginner", Title: "Basic Commands (Beginner)"},
		&cobra.Group{ID: "basic intermediate", Title: "Basic Commands (Intermediate)"},
		&cobra.Group{ID: "deploy", Title: "Deploy Commands"},
		&cobra.Group{ID: "cluster management", Title: "Cluster Management Commands"},
		&cobra.Group{ID: "troubleshooting", Title: "Troubleshooting and Debugging Commands"},
		&cobra.Group{ID: "advanced", Title: "Advanced Commands"},
		&cobra.Group{ID: "settings", Title: "Settings Commands"},
	)

	rootCmd.PersistentFlags().String("as", "", "Username to impersonate for the operation. User could be a regular user or a service account in a namespace.")
	rootCmd.PersistentFlags().StringArray("as-group", []string{}, "Group to impersonate for the operation, this flag can be repeated to specify multiple groups.")
	rootCmd.PersistentFlags().String("as-uid", "", "UID to impersonate for the operation.")
	rootCmd.PersistentFlags().String("cache-dir", "/home/rsteube/.kube/cache", "Default cache directory")
	rootCmd.PersistentFlags().String("certificate-authority", "", "Path to a cert file for the certificate authority")
	rootCmd.PersistentFlags().String("client-certificate", "", "Path to a client certificate file for TLS")
	rootCmd.PersistentFlags().String("client-key", "", "Path to a client key file for TLS")
	rootCmd.PersistentFlags().String("cluster", "", "The name of the kubeconfig cluster to use")
	rootCmd.PersistentFlags().String("context", "", "The name of the kubeconfig context to use")
	rootCmd.PersistentFlags().Bool("disable-compression", false, "If true, opt-out of response compression for all requests to the server")
	rootCmd.Flags().BoolP("help", "h", false, "help for kubectl")
	rootCmd.PersistentFlags().Bool("insecure-skip-tls-verify", false, "If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure")
	rootCmd.PersistentFlags().String("kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")
	rootCmd.PersistentFlags().String("log-flush-frequency", "5s", "Maximum number of seconds between log flushes")
	rootCmd.PersistentFlags().Bool("match-server-version", false, "Require server version to match client version")
	rootCmd.PersistentFlags().StringP("namespace", "n", "", "If present, the namespace scope for this CLI request")
	rootCmd.PersistentFlags().String("password", "", "Password for basic authentication to the API server")
	rootCmd.PersistentFlags().String("profile", "none", "Name of profile to capture. One of (none|cpu|heap|goroutine|threadcreate|block|mutex)")
	rootCmd.PersistentFlags().String("profile-output", "profile.pprof", "Name of the file to write the profile to")
	rootCmd.PersistentFlags().String("request-timeout", "0", "The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests.")
	rootCmd.PersistentFlags().StringP("server", "s", "", "The address and port of the Kubernetes API server")
	rootCmd.PersistentFlags().String("tls-server-name", "", "Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used")
	rootCmd.PersistentFlags().String("token", "", "Bearer token for authentication to the API server")
	rootCmd.PersistentFlags().String("user", "", "The name of the kubeconfig user to use")
	rootCmd.PersistentFlags().String("username", "", "Username for basic authentication to the API server")
	rootCmd.PersistentFlags().StringS("v", "v", "0", "number for the log level verbosity")
	rootCmd.PersistentFlags().String("vmodule", "", "comma-separated list of pattern=N settings for file-filtered logging (only works for the default text log format)")
	rootCmd.PersistentFlags().Bool("warnings-as-errors", false, "Treat warnings received from the server as errors and exit with a non-zero exit code")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir":             carapace.ActionDirectories(),
		"certificate-authority": carapace.ActionFiles(),
		"client-certificate":    carapace.ActionFiles(),
		"client-key":            carapace.ActionFiles(),
		"kubeconfig":            carapace.ActionFiles(),
		"profile":               carapace.ActionValues("none", "cpu", "heap", "goroutine", "threadcreate", "block", "mutex"),
		"profile-output":        carapace.ActionFiles(),
		// TODO add completions for kubeconfig based flags
	})
}
