package finder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaven(t *testing.T) {
	maven, err := FindMaven()
	assert.Nil(t, err)
	assert.NotEmpty(t, maven)
}
