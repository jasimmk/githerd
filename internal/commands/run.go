package commands

import (
	"fmt"
	"os"

	"github.com/careem/githerd/internal/workspaces"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
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
	workspaceName, _ := cmd.Parent().Flags().GetString("workspace")
	fmt.Printf("Running bulk execution in workspace '%s' with command: %s\n", workspaceName, args)
	workspace, err := workspaces.LoadWorkspace(workspaceName)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

}
