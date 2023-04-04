package workspaces

import (
	"fmt"

	"github.com/careem/githerd/pkg/yamlwrapper"
)

func LoadWorkspace(name string) (Workspace, error) {
	workspacePath := GetWorkspacePath(name)
	config := &config{}

	err := yamlwrapper.ReadYAMLFile(workspacePath, config)
	if err != nil {
		return nil, fmt.Errorf("error loading workspace: %s,\nerror: %w", name, err)
	}
	workspace := NewWorkspace(name, workspacePath, config)
	return workspace, nil
}
