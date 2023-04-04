package gitwrapper

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGitWrapper(t *testing.T) {
	// githerd path
	path, err := filepath.Abs("./../../")
	fmt.Println(path)
	gw, err := NewGitWrapper(path)
	assert.NotNil(t, gw)
	assert.Nil(t, err)

	// Test for invalid path
	path = "/non/existent/path"
	gw, err = NewGitWrapper(path)
	assert.Nil(t, gw)
	assert.NotNil(t, err)
}
