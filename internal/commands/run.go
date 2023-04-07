package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/jasimmk/githerd/internal/workspaces"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run bulk execution with arguments in the specified workspace.",
	Long: `Run bulk execution of Git commands with arguments in the specified workspace.
The Git commands and arguments should be provided as a single string argument, e.g. 'git checkout -b test_branch'.
If no workspace is specified, 'default' workspace is used.`,
	Example: "githerd run -w test_workspace 'git checkout -b test_branch'\ngitherd run -w test_workspace run 'git commit -a \"Test commit\"'",
	Run:     runCommand,
}

func runCommand(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}
	workspaceName, _ := cmd.Flags().GetString("workspace")
	fmt.Printf("Running bulk execution in workspace '%s' with command: %s\n", workspaceName, args)
	workspace, err := workspaces.LoadWorkspace(workspaceName)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	// Ensure that all the arguments are provided as a single string, to properly split it as commands
	bulkCmdWithArgs := strings.Split(strings.Join(args[:], " "), " ")
	bulkCmd := bulkCmdWithArgs[0]
	bulkArgs := bulkCmdWithArgs[1:]

	ctx := cmd.Context()
	// Run command in each repository
	for _, repo := range workspace.GetConfig().GetRepositories() {

		fmt.Println(strings.Repeat("=", 120))
		fmt.Printf("Bulk Run: '%s' (%s)\n", repo.Name, repo.Path)
		fmt.Println(strings.Repeat("=", 120))
		gitWrapper, err := repo.GetGitApiWrapper()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", repo.Name, err)
		}
		out, err := gitWrapper.RunCommand(ctx, bulkCmd, bulkArgs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:%s\n", repo.Name, err)
		}
		fmt.Printf("%s", out)
	}

}
