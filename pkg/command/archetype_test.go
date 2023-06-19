package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArchetypeCreate(t *testing.T) {
	create, err := ArchetypeCreate("mvn", "test_data/test-001", "test", "test", "test")
	assert.Nil(t, err)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(create)
}
