package commands

import (
	"fmt"
	"os"

	"github.com/careem/githerd/internal/workspaces"
	"github.com/careem/githerd/pkg/yamlapi"
	"github.com/spf13/cobra"
)

var WorkspaceCmd = &cobra.Command{
	Use:     "workspace",
	Short:   "manage workspaces in githerd",
	Long:    "Run workspace operations on git repositories in the workspace. If no workspace is specified, the 'default' workspace is used.",
	Example: "githerd workspace -w test_workspace init .\ngitherd workspace -w test_workspace show",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

	},
}
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new workspace",
	Long: `Initialize a new workspace at the specified directory/directories. It searches for git repositories in the specified directory and adds them to the workspace.
If no directory is provided, the current directory is used.`,
	Example: "githerd workspace -w test_workspace init .\n",
	Run:     initWorkSpace,
}

var ImportCmd = &cobra.Command{
	Use:     "import",
	Short:   "import a set of repositories, given a file with list of repositories to a workspace",
	Long:    `Clones the repositories specified in file to the directory specified, provided as new line seperated entries. If workspace exists, it adds all the repositories to workspace. If not, it initialize a new workspace at the specified directory`,
	Example: "githerd workspace -w test_workspace import <filename.txt> <directory> .\n",
	Run:     importToWorkSpace,
}

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the contents of a workspace",
	Long: `Show the contents of the specified workspace.
If no workspace is specified, 'default' workspace is used.`,
	Example: "githerd workspace -w test_workspace show",
	Run:     showWorkSpace,
}
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes workspace",
	Long: `Delete workspace specifed.
If no workspace is specified, 'default' workspace is used.`,
	Example: "githerd workspace -w test_workspace delete",
	Run:     deleteWorkSpace,
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
	err = yamlapi.PrintYaml(workspace.GetConfig())
	if err != nil {
		fmt.Printf("Error printing workspace: %s\n", err)
		os.Exit(1)
	}
}

// importToWorkSpace imports a set of repositories, given a file with list of repositories to a workspace
func importToWorkSpace(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Printf("Error: import requires two arguments, <filename.txt> <directory> .\n")
		fmt.Println("")
		cmd.Help()
		os.Exit(1)
	}
	workspaceName, _ := cmd.Parent().Flags().GetString("workspace")
	fmt.Printf("Importing repositories from file '%s' to workspace '%s'...\n", args[0], workspaceName)
	err := workspaces.ImportToWorkspace(cmd.Context(), workspaceName, args[0], args[1])
	if err != nil {
		fmt.Printf("Error importing to workspace: %s\n", err)
		os.Exit(1)
	}
}

// deletwWorkSpace deletes a workspace
func deleteWorkSpace(cmd *cobra.Command, args []string) {
	workspaceName, _ := cmd.Parent().Flags().GetString("workspace")
	fmt.Printf("Deleting workspace '%s'...\n", workspaceName)
	err := workspaces.DeleteWorkspace(workspaceName)
	if err != nil {
		fmt.Printf("Error deleting workspace: %s\n", err)
		os.Exit(1)
	}
}
