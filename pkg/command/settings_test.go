package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocalRepository(t *testing.T) {
	repoDirectory, err := GetLocalRepositoryDirectory("D:\\soft\\apache-maven-3.8.5\\bin\\mvn")
	assert.Nil(t, err)
	assert.NotEmpty(t, repoDirectory)
	t.Log(repoDirectory)
}
