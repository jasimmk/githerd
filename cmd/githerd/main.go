package main

import (
	"context"
	"fmt"
	"os"

	"github.com/careem/githerd/internal/commands"
	"github.com/careem/githerd/internal/constants"
	"github.com/careem/githerd/pkg/file"
	"github.com/spf13/viper"
)

var Version = "v0.0.0"
var Commit = "000"
var TagCommit = ""

func main() {
	var configFile string
	var workspace string

	rootCmd := commands.RootCmd
	versionCmd := commands.VersionCmd

	WorkspaceCmd := commands.WorkspaceCmd
	initCmd := commands.InitCmd
	showCmd := commands.ShowCmd

	RunCmd := commands.RunCmd

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "~/.githerd/config.yaml", "config file")
	RunCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "default", "Specify the workspace to use, By default 'default' workspace is used")
	WorkspaceCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "default", "Specify the workspace to use, By default 'default' workspace is used")
	// bulkCmd.PersistentFlags().StringVarP(&workspace, "profile", "p", "default", "Specify the profile to use for the workspace, By default 'default' profile is used")
	WorkspaceCmd.AddCommand(initCmd, showCmd)
	rootCmd.AddCommand(WorkspaceCmd, RunCmd, versionCmd)

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
	configFile, err = file.AbsPath(configFile)
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
