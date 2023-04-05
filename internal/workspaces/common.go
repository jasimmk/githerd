package workspaces

import (
	"fmt"

	"github.com/careem/githerd/internal/constants"
	"github.com/careem/githerd/pkg/file"
)

func GetWorkspacePath(name string) string {
	path, err := file.AbsPath(constants.GLOBAL_WORKSPACE_DIR)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s/%s.yaml", path, name)
}
