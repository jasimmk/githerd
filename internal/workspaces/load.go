package workspaces

import (
	"fmt"

	"github.com/jasimmk/githerd/pkg/yamlapi"
)

func LoadWorkspace(name string) (Workspace, error) {
	workspacePath := GetWorkspacePath(name)
	config := &config{}

	err := yamlapi.ReadYAMLFile(workspacePath, config)
	if err != nil {
		return nil, fmt.Errorf("error loading workspace: %s,\nerror: %w", name, err)
	}
	workspace := NewWorkspace(name, workspacePath, config)
	return workspace, nil
}
