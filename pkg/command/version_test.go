package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	version, err := Version("D:\\soft\\apache-maven-3.8.5\\bin\\mvn")
	assert.Nil(t, err)
	assert.NotEmpty(t, version)
	t.Log(version)
}
