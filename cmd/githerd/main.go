package main

import (
	"context"
	"fmt"
	"os"

	"github.com/careem/githerd/internal/commands"
	"github.com/careem/githerd/internal/constants"
	"github.com/careem/githerd/pkg/filewrapper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Version = "v0.0.0"
var Commit = "000"
var TagCommit = ""

var workspaceCmd = &cobra.Command{
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

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run bulk execution with arguments in the specified workspace.",
	Long: `Run bulk execution of Git commands with arguments in the specified workspace.
The Git commands and arguments should be provided as a single string argument, e.g. 'git checkout -b test_branch'.
If no workspace is specified, 'default' workspace is used.`,
	Example: "githerd run -w test_workspace 'git checkout -b test_branch'\ngitherd run -w test_workspace run 'git commit -a \"Test commit\"'",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
		workspace, _ := cmd.Parent().Flags().GetString("workspace")
		fmt.Printf("Running bulk execution in workspace '%s' with command: %s\n", workspace, args)
		// TODO: Implement bulk execution logic
	},
}

func main() {
	var configFile string
	var workspace string

	rootCmd := commands.RootCmd
	versionCmd := commands.VersionCmd

	initCmd := commands.InitCmd
	showCmd := commands.ShowCmd

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "~/.githerd/config.yaml", "config file")
	runCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "default", "Specify the workspace to use, By default 'default' workspace is used")
	workspaceCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "default", "Specify the workspace to use, By default 'default' workspace is used")
	// bulkCmd.PersistentFlags().StringVarP(&workspace, "profile", "p", "default", "Specify the profile to use for the workspace, By default 'default' profile is used")
	workspaceCmd.AddCommand(initCmd, showCmd)
	rootCmd.AddCommand(workspaceCmd, runCmd, versionCmd)

	// reading viper config
	readConfig(configFile)

	// Set contexts
	ctx := context.Background()
	ctx = setVersion(ctx)

	// Run the command
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func readConfig(configFile string) {
	var err error
	configFile, err = filewrapper.AbsPath(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if _, err := os.Stat(configFile); err != nil {
		fmt.Printf("missing config file %s \n. Please check the README.md of the project\n", configFile)
		os.Exit(1)
	}
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// setVersion sets the version of the application in the context
func setVersion(ctx context.Context) context.Context {
	version := Version
	if version == "" {
		version = "v0.0.0"
	}
	if Commit != "" {
		if TagCommit != Commit {
			version = fmt.Sprintf("%s-%s", Version, Commit)
		}
	}
	return context.WithValue(ctx, constants.CtxKey("Version"), version)
}
