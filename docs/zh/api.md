# API 参考

本文档提供了 Maven SDK Go 的详细 API 参考。

## 包概览

### Finder

Finder 包用于在 Maven 本地仓库中查找 JAR 文件。

```go
package finder

// FindJar 在本地仓库中查找指定坐标的 JAR 文件
func FindJar(groupId, artifactId, version string) (string, error)

// FindJarWithClassifier 查找带分类器的 JAR 文件
func FindJarWithClassifier(groupId, artifactId, version, classifier string) (string, error)
```

### Command

Command 包提供了执行 Maven 命令的功能。

```go
package command

// NewMavenCommand 创建一个新的 Maven 命令实例
func NewMavenCommand() *MavenCommand

// Execute 执行 Maven 命令
func (cmd *MavenCommand) Execute(args ...string) error

// SetWorkingDirectory 设置工作目录
func (cmd *MavenCommand) SetWorkingDirectory(dir string)
```

### Local Repository

Local Repository 包用于解析和管理 Maven 本地仓库。

```go
package local_repository

// GetLocalRepositoryPath 获取本地仓库路径
func GetLocalRepositoryPath() (string, error)

// ParseArtifactPath 解析构件路径
func ParseArtifactPath(groupId, artifactId, version string) string
```

### Installer

Installer 包提供了 Maven 的自动安装功能。

```go
package installer

// InstallMaven 自动安装指定版本的 Maven
func InstallMaven(version string) error

// GetInstalledMavenVersion 获取已安装的 Maven 版本
func GetInstalledMavenVersion() (string, error)
```

## 使用示例

### 查找 JAR 文件

```go
jarPath, err := finder.FindJar("org.springframework", "spring-core", "5.3.21")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Spring Core JAR: %s\n", jarPath)
```

### 执行 Maven 命令

```go
cmd := command.NewMavenCommand()
cmd.SetWorkingDirectory("/path/to/project")
err := cmd.Execute("clean", "install")
if err != nil {
    log.Fatal(err)
}
```

### 获取本地仓库路径

```go
repoPath, err := local_repository.GetLocalRepositoryPath()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("本地仓库路径: %s\n", repoPath)
```

### 安装 Maven

```go
err := installer.InstallMaven("3.8.6")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Maven 安装成功")
```