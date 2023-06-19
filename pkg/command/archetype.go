package command

// ArchetypeCreate 创建项目
// executable: mvn可执行文件的位置
// directory: 生成的项目的位置
// groupId / artifactId /
func ArchetypeCreate(executable string, directory string, groupId, artifactId, version string) (string, error) {
	// maven3.0.5以上版本舍弃了create，使用generate生成项目
	// mvn archetype:create -DgroupId=org.sonatype.mavenbook.ch03 -DartifactId=simple -DpackageName=org.sonatype.mavenbook
	return ExecForStdout(executable, "archetype:generate", "-DoutputDirectory="+directory, "-DgroupId="+groupId, "-DartifactId="+artifactId, "-DarchetypeVersion="+version, "-DinteractiveMode=false")
}
