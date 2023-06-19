package local_repository

import (
	"github.com/scagogogo/mvn-sdk/pkg/checker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseLocalRepositoryDirectory(t *testing.T) {
	maven, err := finder.FindMaven()
	assert.Nil(t, err)
	directory := ParseLocalRepositoryDirectory(maven)
	assert.NotEmpty(t, directory)
}

func TestFindDirectory(t *testing.T) {
	maven, err := finder.FindMaven()
	assert.Nil(t, err)
	directory := ParseLocalRepositoryDirectory(maven)
	findDirectory, err := FindDirectory(directory, "joda-time", "joda-time", "2.10.10")
	assert.Nil(t, err)
	assert.NotEmpty(t, findDirectory)
}

func TestFindJar(t *testing.T) {
	maven, err := finder.FindMaven()
	assert.Nil(t, err)
	directory := ParseLocalRepositoryDirectory(maven)
	findDirectory, err := FindJar(directory, "joda-time", "joda-time", "2.10.10")
	assert.Nil(t, err)
	assert.NotEmpty(t, findDirectory)
}