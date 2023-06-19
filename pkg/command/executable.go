package command

import "path/filepath"

// BuildExecutable 根据 %MAVEN_HOME% 构建mvn可执行文件的路径
func BuildExecutable(mavenHomeDirectory string) string {
	return filepath.Join(mavenHomeDirectory, "bin/mvn")
}
