package commands

import (
	"fmt"
	"os"

	"github.com/jasimmk/githerd/pkg/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SetupCommands sets up the commands and returns the root command
func SetupCommands() *cobra.Command {
	var configFile string
	var workspace string

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "~/.githerd/config.yaml", "config file")
	runCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "default", "Specify the workspace to use, By default 'default' workspace is used")
	workspaceCmd.PersistentFlags().StringVarP(&workspace, "workspace", "w", "default", "Specify the workspace to use, By default 'default' workspace is used")
	// bulkCmd.PersistentFlags().StringVarP(&workspace, "profile", "p", "default", "Specify the profile to use for the workspace, By default 'default' profile is used")
	workspaceCmd.AddCommand(initCmd, showCmd, importCmd, deleteCmd)
	rootCmd.AddCommand(workspaceCmd, runCmd, versionCmd)

	// reading viper config
	readConfig(configFile)
	return rootCmd
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
