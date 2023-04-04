package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/careem/githerd/internal/constants"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Short: "A command line tool for bulk execution of Git commands in a workspace",
	Long: `githerd is a command line tool that allows for bulk execution of Git commands in a workspace.
	It supports initializing a new workspace, showing the contents of a workspace, and running bulk Git commands with arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		version := getVersion(ctx)
		fmt.Printf("Welcome to githerd!\nVersion: %s\n", version)
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}
	},
}

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  "Prints current version of the application",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		version := getVersion(ctx)
		fmt.Printf("githerd version: %s\n", version)
	},
}

func getVersion(ctx context.Context) string {
	ver := constants.CtxKey("Version")
	return ctx.Value(ver).(string)

}
