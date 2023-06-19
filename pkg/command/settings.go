package command

// GetLocalRepositoryDirectory 从命令获取本地仓库的位置
func GetLocalRepositoryDirectory(executable string) (string, error) {
	return ExecForStdout(executable, "help:evaluate", "-Dexpression=settings.localRepository", "-q", "-DforceStdout")
}
