package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jasimmk/githerd/internal/commands"
	"github.com/jasimmk/githerd/internal/constants"
)

var Version = "v0.0.0"
var Commit = "000"
var TagCommit = ""

func main() {
	rootCmd := commands.SetupCommands()
	// Set contexts
	ctx := context.Background()
	ctx = setVersion(ctx)

	// Run the command
	if err := rootCmd.ExecuteContext(ctx); err != nil {
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
