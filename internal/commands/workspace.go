package commands

import (
	"fmt"
	"os"

	"github.com/careem/githerd/internal/workspaces"
	"github.com/careem/githerd/pkg/yamlwrapper"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new workspace",
	Long: `Initialize a new workspace at the specified directory/directories. It searches for git repositories in the specified directory and adds them to the workspace.
If no directory is provided, the current directory is used.`,
	Example: "githerd workspace -w test_workspace init .\n",
	Run:     initWorkSpace,
}

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the contents of a workspace",
	Long: `Show the contents of the specified workspace.
If no workspace is specified, 'default' workspace is used.`,
	Example: "githerd workspace -w test_workspace show",
	Run:     showWorkSpace,
}

// create a new workspace
func initWorkSpace(cmd *cobra.Command, args []string) {
	dir := []string{"."}
	if len(args) > 0 {
		dir = args
	}
	workspaceName, _ := cmd.Parent().Flags().GetString("workspace")
	fmt.Printf("Initializing workspace '%s' with directories: %s...\n", workspaceName, dir)
	err := workspaces.CreateWorkspace(cmd.Context(), workspaceName, dir)
	if err != nil {
		fmt.Printf("Error initializing workspace: %s\n", err)
		os.Exit(1)
	}
}

// show the contents of a workspace
func showWorkSpace(cmd *cobra.Command, args []string) {
	workspaceName, _ := cmd.Parent().Flags().GetString("workspace")
	fmt.Printf("Showing contents of workspace '%s'...\n", workspaceName)
	workspace, err := workspaces.LoadWorkspace(workspaceName)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	err = yamlwrapper.PrintYaml(workspace.GetConfig())
	if err != nil {
		fmt.Printf("Error printing workspace: %s\n", err)
		os.Exit(1)
	}
}
