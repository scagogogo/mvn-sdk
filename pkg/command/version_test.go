package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	// 使用系统PATH中的mvn
	executable := "mvn"
	version, err := Version(executable)
	assert.Nil(t, err)
	assert.NotEmpty(t, version)
	t.Log(version)
}
