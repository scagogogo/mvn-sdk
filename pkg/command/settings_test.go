package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalRepository(t *testing.T) {
	// 使用系统PATH中的mvn，或者使用环境变量
	executable := "mvn"
	repoDirectory, err := GetLocalRepositoryDirectory(executable)
	assert.Nil(t, err)
	assert.NotEmpty(t, repoDirectory)
	t.Log(repoDirectory)
}
