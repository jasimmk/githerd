package workspaces

import "os"

// DeleteWorkspace deletes the workspace with the given name.
func DeleteWorkspace(name string) error {
	workspacePath := GetWorkspacePath(name)

	err := os.Remove(workspacePath)
	return err

}
