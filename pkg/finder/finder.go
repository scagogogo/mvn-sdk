package finder

import (
	"errors"
	"github.com/scagogogo/mvn-sdk/pkg/command"
	"os"
	"strings"
)

// ErrNotFoundMaven 没有找到Maven
var ErrNotFoundMaven = errors.New("not found maven")

// FindMaven 查找本机已经安装的Maven
func FindMaven() (string, error) {

	// 尝试从PATH中找到可以执行的Maven
	stdout, err := command.ExecForStdout("mvn", "--help")
	if err == nil && strings.Contains(stdout, "usage: mvn [options] [<goal(s)>] [<phase(s)>]") {
		return "mvn", nil
	}

	// 再从几个环境变量中查找
	envNameSlice := []string{"M2_HOME", "MAVEN_HOME"}
	for _, envName := range envNameSlice {
		getenv := os.Getenv(envName)
		if envName == "" {
			continue
		}
		if Check(getenv) {
			return command.BuildExecutable(getenv), nil
		}
	}

	return "", ErrNotFoundMaven
}

// Check 检查是否是合法的Maven目录，依据是是否有mvn的可执行文件
func Check(mavenHomeDirectory string) bool {
	executable := command.BuildExecutable(mavenHomeDirectory)
	stat, err := os.Stat(executable)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}
