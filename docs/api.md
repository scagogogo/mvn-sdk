# API Reference

本文档提供了 Go Maven SDK 的完整 API 参考。

## Packages

### Finder {#finder}

Maven 查找器相关功能，用于在系统上查找已安装的 Maven。

#### Functions

##### `FindMaven() (string, error)`

查找本地已安装的 Maven 可执行文件。

**返回值:**
- `string`: Maven 可执行文件的完整路径
- `error`: 如果未找到 Maven 则返回错误

**示例:**
```go
maven, err := finder.FindMaven()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found Maven at: %s\n", maven)
```

##### `Check(mavenHomeDirectory string) bool`

检查指定路径是否为合法的 Maven 安装目录。

**参数:**
- `mavenHomeDirectory`: Maven 安装目录路径

**返回值:**
- `bool`: 如果是合法的 Maven 目录则返回 true

### Command {#command}

Maven 命令执行相关功能，提供执行 Maven 命令的接口。

#### Types

##### `Options`

命令执行选项结构体：

```go
type Options struct {
    Executable string    // Maven 可执行文件路径
    Args       []string  // 命令参数
    Stdin      io.Reader // 标准输入
    Stdout     io.Writer // 标准输出
    Stderr     io.Writer // 标准错误
}
```

#### Functions

##### `Exec(options *Options) error`

执行 Maven 命令。

**参数:**
- `options`: 命令执行选项

##### `ExecForStdout(executable string, args ...string) (string, error)`

执行 Maven 命令并返回标准输出。

**参数:**
- `executable`: Maven 可执行文件路径
- `args`: 命令参数列表

**返回值:**
- `string`: 命令的标准输出
- `error`: 错误信息

##### `DependencyGet(executable, groupId, artifactId, version string) (string, error)`

下载指定的 Maven 依赖。

**参数:**
- `executable`: Maven 可执行文件路径
- `groupId`: Maven 组 ID
- `artifactId`: Maven 工件 ID
- `version`: 版本号

##### `Version(executable string) (string, error)`

获取 Maven 版本信息。

**参数:**
- `executable`: Maven 可执行文件路径

##### `GetLocalRepositoryDirectory(executable string) (string, error)`

获取 Maven 本地仓库目录路径。

**参数:**
- `executable`: Maven 可执行文件路径

##### `ArchetypeCreate(executable string, directory, groupId, artifactId, version string) (string, error)`

使用 Maven archetype 创建新项目。

**参数:**
- `executable`: Maven 可执行文件路径
- `directory`: 项目创建目录
- `groupId`: 项目组 ID
- `artifactId`: 项目工件 ID
- `version`: 项目版本

### Local Repository {#local-repository}

Maven 本地仓库管理功能，用于操作 Maven 本地仓库。

#### Variables

##### `DefaultLocalRepositoryDirectory string`

默认的 Maven 本地仓库目录路径（通常是 `~/.m2/repository`）。

#### Functions

##### `ParseLocalRepositoryDirectory(executable string) string`

解析获取 Maven 本地仓库目录路径。

**参数:**
- `executable`: Maven 可执行文件路径

**返回值:**
- `string`: 本地仓库目录路径

##### `BuildDirectory(groupId, artifactId, version string) string`

根据 GAV 坐标构建本地仓库目录路径。

**参数:**
- `groupId`: 组 ID
- `artifactId`: 工件 ID
- `version`: 版本号

**返回值:**
- `string`: 构建的目录路径

##### `FindDirectory(localRepositoryDirectory, groupId, artifactId, version string) (string, error)`

在本地仓库中查找指定 GAV 坐标的目录。

**参数:**
- `localRepositoryDirectory`: 本地仓库根目录
- `groupId`: 组 ID
- `artifactId`: 工件 ID
- `version`: 版本号

**返回值:**
- `string`: 找到的目录路径
- `error`: 错误信息

##### `FindJar(localRepositoryDirectory, groupId, artifactId, version string) (string, error)`

在本地仓库中查找指定 GAV 坐标的 JAR 文件。

**参数:**
- `localRepositoryDirectory`: 本地仓库根目录
- `groupId`: 组 ID
- `artifactId`: 工件 ID
- `version`: 版本号

**返回值:**
- `string`: JAR 文件的完整路径
- `error`: 错误信息

**示例:**
```go
jarPath, err := local_repository.FindJar(
    "/home/user/.m2/repository", 
    "org.apache.commons", 
    "commons-lang3", 
    "3.12.0",
)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Found JAR at: %s\n", jarPath)
```

### Installer {#installer}

Maven 自动安装功能，支持在不同操作系统上自动安装 Maven。

#### Functions

##### `Install() (string, error)`

根据当前操作系统自动安装 Maven。

**返回值:**
- `string`: 安装的 Maven 可执行文件路径
- `error`: 安装过程中的错误

##### `InstallWindows() (string, error)`

在 Windows 系统上安装 Maven。

##### `InstallLinux() (string, error)`

在 Linux 系统上安装 Maven。

##### `InstallMacOS() (string, error)`

在 macOS 系统上安装 Maven。

## 错误处理

### 常见错误类型

- `finder.ErrNotFoundMaven`: 未找到 Maven 安装
- 文件不存在错误: 指定的 JAR 文件或目录不存在
- 命令执行错误: Maven 命令执行失败

### 错误处理示例

```go
maven, err := finder.FindMaven()
if err != nil {
    if errors.Is(err, finder.ErrNotFoundMaven) {
        fmt.Println("Maven not found. Attempting to install...")
        maven, err = installer.Install()
        if err != nil {
            log.Fatal("Failed to install Maven:", err)
        }
    } else {
        log.Fatal("Error finding Maven:", err)
    }
}
```