# Go Maven SDK API Documentation

## Overview

Go Maven SDK 是一个用于在 Go 语言中操作 Maven 的工具库。它提供了查找 Maven、管理本地仓库、执行 Maven 命令等功能。

## Packages

### pkg/finder

Maven 查找器相关功能。

#### Functions

##### `FindMaven() (string, error)`

查找本地已安装的 Maven 可执行文件。

**Returns:**
- `string`: Maven 可执行文件路径
- `error`: 错误信息，如果未找到 Maven

**Example:**
```go
maven, err := finder.FindMaven()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found Maven at: %s\n", maven)
```

##### `Check(mavenHomeDirectory string) bool`

检查指定路径是否为合法的 Maven 目录。

**Parameters:**
- `mavenDirectory`: Maven 安装目录路径

**Returns:**
- `bool`: 是否为合法的 Maven 目录

### pkg/command

Maven 命令执行相关功能。

#### Types

##### `Options`

命令执行选项结构体。

```go
type Options struct {
    Executable string
    Args       []string
    Stdin      io.Reader
    Stdout     io.Writer
    Stderr     io.Writer
}
```

#### Functions

##### `Exec(options *Options) error`

执行 Maven 命令。

**Parameters:**
- `options`: 命令执行选项

##### `ExecForStdout(executable string, args ...string) (string, error)`

执行 Maven 命令并返回标准输出。

**Parameters:**
- `executable`: Maven 可执行文件路径
- `args`: 命令参数

**Returns:**
- `string`: 标准输出内容
- `error`: 错误信息

**Example:**
```go
output, err := command.ExecForStdout("mvn", "-version")
if err != nil {
    log.Fatal(err)
}
fmt.Println(output)
```

##### `DependencyGet(executable, groupId, artifactId, version string) (string, error)`

下载指定的 Maven 依赖。

**Parameters:**
- `executable`: Maven 可执行文件路径
- `groupId`: 组 ID
- `artifactId`: 工件 ID
- `version`: 版本号

**Returns:**
- `string`: 命令输出
- `error`: 错误信息

##### `Version(executable string) (string, error)`

获取 Maven 版本信息。

**Parameters:**
- `executable`: Maven 可执行文件路径

##### `GetLocalRepositoryDirectory(executable string) (string, error)`

获取本地仓库目录路径。

##### `ArchetypeCreate(executable string, directory, groupId, artifactId, version string) (string, error)`

使用 archetype 创建 Maven 项目。

### pkg/local_repository

Maven 本地仓库管理功能。

#### Variables

##### `DefaultLocalRepositoryDirectory string`

默认本地仓库目录路径。

#### Functions

##### `ParseLocalRepositoryDirectory(executable string) string`

解析本地仓库目录路径。

**Parameters:**
- `executable`: Maven 可执行文件路径

**Returns:**
- `string`: 本地仓库目录路径

##### `BuildDirectory(groupId, artifactId, version string) string`

构建 GAV 坐标的目录路径。

**Parameters:**
- `groupId`: 组 ID
- `artifactId`: 工件 ID
- `version`: 版本号

**Returns:**
- `string`: GAV 目录路径

##### `FindDirectory(localRepositoryDirectory, groupId, artifactId, version string) (string, error)`

在本地仓库中查找指定 GAV 的目录。

**Parameters:**
- `localRepositoryDirectory`: 本地仓库目录路径
- `groupId`: 组 ID
- `artifactId`: 工件 ID
- `version`: 版本号

**Returns:**
- `string`: GAV 目录路径
- `error`: 错误信息

##### `FindJar(localRepositoryDirectory, groupId, artifactId, version string) (string, error)`

在本地仓库中查找指定 GAV 的 JAR 文件。

**Parameters:**
- `localRepositoryDirectory`: 本地仓库目录路径
- `groupId`: 组 ID
- `artifactId`: 工件 ID
- `version`: 版本号

**Returns:**
- `string`: JAR 文件路径
- `error`: 错误信息

**Example:**
```go
jarPath, err := local_repository.FindJar("/home/user/.m2/repository", "org.apache.commons", "commons-lang3", "3.12.0")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found JAR at: %s\n", jarPath)
```

### pkg/installer

Maven 自动安装功能。

#### Functions

##### `Install() (string, error)`

根据当前操作系统自动安装 Maven。

**Returns:**
- `string`: Maven 安装路径
- `error`: 错误信息

##### `InstallWindows() (string, error)`

在 Windows 系统上安装 Maven。

##### `InstallLinux() (string, error)`

在 Linux 系统上安装 Maven。

##### `InstallMacOS() (string, error)`

在 macOS 系统上安装 Maven。

## Examples

### 基本用法

```go
package main

import (
    "fmt"
    "github.com/scagogogo/mvn-sdk/pkg/finder"
    "github.com/scagogogo/mvn-sdk/pkg/local_repository"
    "github.com/scagogogo/mvn-sdk/pkg/command"
)

func main() {
    // 查找 Maven
    maven, err := finder.FindMaven()
    if err != nil {
        panic(err)
    }
    
    // 获取本地仓库路径
    repoDir := local_repository.ParseLocalRepositoryDirectory(maven)
    
    // 查找 JAR 文件
    jarPath, err := local_repository.FindJar(repoDir, "org.apache.commons", "commons-lang3", "3.12.0")
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("JAR path: %s\n", jarPath)
    
    // 下载依赖
    output, err := command.DependencyGet(maven, "org.apache.commons", "commons-lang3", "3.12.0")
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Download output: %s\n", output)
}
```

### 错误处理

```go
maven, err := finder.FindMaven()
if err != nil {
    if errors.Is(err, finder.ErrNotFoundMaven) {
        fmt.Println("Maven not found. Please install Maven first.")
    } else {
        fmt.Printf("Error finding Maven: %v\n", err)
    }
    os.Exit(1)
}
```

## Error Handling

### Common Errors

- `finder.ErrNotFoundMaven`: 未找到 Maven 安装
- File not found errors: 指定的 JAR 文件或目录不存在
- Command execution errors: Maven 命令执行失败

## Contributing

欢迎提交 Issue 和 Pull Request 来改进这个项目。

## License

[License information]