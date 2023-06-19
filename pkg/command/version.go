package command

// Version 获取Maven的版本
func Version(executable string) (string, error) {
	return ExecForStdout(executable, "-v")
}
