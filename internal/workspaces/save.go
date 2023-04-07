package workspaces

import "github.com/jasimmk/githerd/pkg/yamlapi"

func Save(name string, workspaceConfig Config) error {
	workspacePath := GetWorkspacePath(name)
	return yamlapi.WriteYamlFile(workspacePath, workspaceConfig)

}
