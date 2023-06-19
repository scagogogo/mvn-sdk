package local_repository

import (
	"errors"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/scagogogo/mvn-sdk/pkg/command"
	"os"
	"path/filepath"
	"strings"
)

// DefaultLocalRepositoryDirectory 默认的仓库位置，${user.home}/.m2/repository
// Windows 7：C:/Documents and Settings/<用户名>/.m2/repository
// Windows 10：C:/Users/<用户名>/.m2/repository
// Linux：/home/<用户名>/.m2/repository
// Mac：/Users/<用户名>/.m2/repository
var DefaultLocalRepositoryDirectory string

func init() {
	dir, err := homedir.Dir()
	if err != nil {
		return
	}
	DefaultLocalRepositoryDirectory = filepath.Join(dir, ".m2/repository/")
}

// ParseLocalRepositoryDirectory 解析本地仓库的位置
func ParseLocalRepositoryDirectory(executable string) string {

	// 尝试从已经安装的Maven中寻找仓库的位置
	directory, err := command.GetLocalRepositoryDirectory(executable)
	if err == nil && directory != "" {
		return directory
	}

	// 找不到则返回默认的仓库位置
	return DefaultLocalRepositoryDirectory
}

// BuildDirectory 构造GAV的相对路径
func BuildDirectory(groupId, artifactId, version string) string {
	return filepath.Join(strings.ReplaceAll(groupId, ".", "/"), artifactId, version)
}

// FindDirectory 在本地仓库中寻找GAV
func FindDirectory(localRepositoryDirectory string, groupId, artifactId, version string) (string, error) {
	gavDirectory := filepath.Join(localRepositoryDirectory, BuildDirectory(groupId, artifactId, version))
	stat, err := os.Stat(gavDirectory)
	if err != nil {
		return "", err
	}
	if stat.IsDir() {
		return gavDirectory, nil
	} else {
		return "", errors.New("not a directory")
	}
}

// FindJar 在本地仓库中寻找给定的GAV的Jar包
func FindJar(localRepositoryDirectory string, groupId, artifactId, version string) (string, error) {
	directory, err := FindDirectory(localRepositoryDirectory, groupId, artifactId, version)
	if err != nil {
		return "", err
	}
	jarPath := filepath.Join(directory, fmt.Sprintf("%s-%s.jar", artifactId, version))
	stat, err := os.Stat(jarPath)
	if err != nil {
		return "", err
	}
	if !stat.IsDir() {
		return jarPath, nil
	} else {
		return "", errors.New("is directory, need a file")
	}
}
