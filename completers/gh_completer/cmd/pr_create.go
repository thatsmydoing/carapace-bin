package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a pull request",
	GroupID: "general",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_createCmd).Standalone()

	pr_createCmd.Flags().StringSliceP("assignee", "a", []string{}, "Assign people by their `login`. Use \"@me\" to self-assign.")
	pr_createCmd.Flags().StringP("base", "B", "", "The `branch` into which you want your code merged")
	pr_createCmd.Flags().StringP("body", "b", "", "Body for the pull request")
	pr_createCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	pr_createCmd.Flags().BoolP("draft", "d", false, "Mark pull request as a draft")
	pr_createCmd.Flags().BoolP("fill", "f", false, "Do not prompt for title/body and just use commit info")
	pr_createCmd.Flags().StringP("head", "H", "", "The `branch` that contains commits for your pull request (default: current branch)")
	pr_createCmd.Flags().StringSliceP("label", "l", []string{}, "Add labels by `name`")
	pr_createCmd.Flags().StringP("milestone", "m", "", "Add the pull request to a milestone by `name`")
	pr_createCmd.Flags().Bool("no-maintainer-edit", false, "Disable maintainer's ability to modify pull request")
	pr_createCmd.Flags().StringSliceP("project", "p", []string{}, "Add the pull request to projects by `name`")
	pr_createCmd.Flags().String("recover", "", "Recover input from a failed run of create")
	pr_createCmd.Flags().StringSliceP("reviewer", "r", []string{}, "Request reviews from people or teams by their `handle`")
	pr_createCmd.Flags().StringP("title", "t", "", "Title for the pull request")
	pr_createCmd.Flags().BoolP("web", "w", false, "Open the web browser to create a pull request")
	prCmd.AddCommand(pr_createCmd)

	carapace.Gen(pr_createCmd).FlagCompletion(carapace.ActionMap{
		"assignee":  action.ActionAssignableUsers(pr_createCmd).UniqueList(","),
		"base":      action.ActionBranches(pr_createCmd),
		"body":      action.ActionBody(pr_createCmd),
		"body-file": carapace.ActionFiles(),
		"head":      action.ActionBranches(pr_createCmd),
		"label":     action.ActionLabels(pr_createCmd).UniqueList(","),
		"milestone": action.ActionMilestones(pr_createCmd),
		"project":   action.ActionProjects(pr_createCmd, action.ProjectOpts{Open: true}),
		"reviewer":  action.ActionAssignableUsers(pr_createCmd).UniqueList(","),
	})
}
