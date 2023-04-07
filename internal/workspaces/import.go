package workspaces

import (
	"context"
	"fmt"

	"github.com/jasimmk/githerd/pkg/execapi"
	"github.com/jasimmk/githerd/pkg/file"
)

// ImportToWorkspace imports a set of repositories, given a file with list of repositories to a workspace
func ImportToWorkspace(ctx context.Context, name string, filePath string, dir string) error {

	// File and Dir management
	filPath, err := file.AbsPath(filePath)
	if err != nil {
		return err
	}
	repoUris, err := file.ReadLines(filPath, true)
	if err != nil {
		return err
	}

	dirPath, err := file.AbsPath(dir)
	if err != nil {
		return err
	}
	err = file.CreateDirIfNotExists(dirPath)
	if err != nil {
		return err
	}
	// Clone the repositories to the directory
	for _, repoUri := range repoUris {
		fmt.Println("Cloning repository: ", repoUri)
		//TODO: Fix with proper configuration
		// _, err := gitapi.CloneRepository(repoUri, dirPath, "")
		out, err := execapi.RunShellExec(fmt.Sprintf("git clone %s", repoUri), dirPath)
		if err != nil {
			fmt.Printf("Cloning repository:%s failed: %s\nOutput: %s", repoUri, err, out)
		}
		fmt.Println(string(out))
	}
	UpsertWorkspace(ctx, name, []string{dirPath})
	return err
}
