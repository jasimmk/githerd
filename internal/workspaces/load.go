package workspaces

import (
	"fmt"

	"github.com/careem/githerd/pkg/yamlwrapper"
)

func LoadWorkspace(name string) (Workspace, error) {
	workspacePath := GetWorkspacePath(name)
	config := WorkspaceConfig{}

	err := yamlwrapper.ReadYAMLFile(workspacePath, config)
	if err != nil {
		return nil, fmt.Errorf("error loading workspace: %s, error: %w", workspacePath, err)
	}
	workspace := NewWorkspace(name, workspacePath, config)
	return workspace, nil
}
