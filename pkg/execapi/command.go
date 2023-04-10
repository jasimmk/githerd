package execapi

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-cmd/cmd"
)

// RunShellExec runs a command with shell and returns the output and error
func RunShellExec(ctx context.Context, wg *sync.WaitGroup, command, workDir string) {
	/*
		cmd := exec.Command("sh", "-c", command)
		cmd.Dir = workDir
		return cmd.CombinedOutput()
	*/
	// Runs command with output in async mode
	runCmd := cmd.NewCmd("sh", "-c", command)
	runCmd.Dir = workDir

	var statusChan <-chan cmd.Status = runCmd.Start()

	// Print last line of stdout every 2s
	var ticker *time.Ticker = time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			_ = runCmd.Status()

		}
	}()
	var cleanup = func() {
		ticker.Stop()
		wg.Done()
	}
	// TODO: Configure timeout
	timeout := time.NewTimer(5 * time.Minute)
	for {
		select {
		case finalStatus := <-statusChan:
			if finalStatus.Exit == 0 {
				fmt.Printf("Command:%s finished successfully:\n%v\n", command, strings.Join(finalStatus.Stdout, "\n"))
			} else {
				fmt.Printf("Command:%s finished with error:\n%v\n", command, strings.Join(finalStatus.Stderr, "\n"))
			}
			cleanup()
			return
		case <-timeout.C:
			runCmd.Stop()
			fmt.Printf("Command:%s timed out", command)
			cleanup()
			return
		}
	}
}
