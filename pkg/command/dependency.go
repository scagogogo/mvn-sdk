package command

// 可以修改删除格式类型

// DependencyTree 打印出整个项目的依赖树，需要现在处于项目的根目录下
func DependencyTree() {
	// mvn dependency:tree

}

// DependencyGet 下载指定的包到本地
func DependencyGet(executable string, groupId, artifactId, version string) (string, error) {
	// mvn dependency:get -DgroupId=<groupId> -DartifactId=<artifactId> -Dversion=<version>
	return ExecForStdout(executable, "dependency:get", "-DgroupId="+groupId, "-DartifactId="+artifactId, "-Dversion="+version)
}
