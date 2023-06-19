package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDependencyGet(t *testing.T) {
	stdout, err := DependencyGet("", "joda-time", "joda-time", "2.10.10")
	assert.Nil(t, err)
	assert.NotEmpty(t, stdout)
}
