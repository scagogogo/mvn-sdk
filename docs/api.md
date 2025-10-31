# API Reference

This document provides detailed API reference for Maven SDK Go.

## Package Overview

### Finder

The Finder package is used to find JAR files in the Maven local repository.

```go
package finder

// FindJar finds a JAR file with the specified coordinates in the local repository
func FindJar(groupId, artifactId, version string) (string, error)

// FindJarWithClassifier finds a JAR file with classifier
func FindJarWithClassifier(groupId, artifactId, version, classifier string) (string, error)
```

### Command

The Command package provides functionality to execute Maven commands.

```go
package command

// NewMavenCommand creates a new Maven command instance
func NewMavenCommand() *MavenCommand

// Execute executes a Maven command
func (cmd *MavenCommand) Execute(args ...string) error

// SetWorkingDirectory sets the working directory
func (cmd *MavenCommand) SetWorkingDirectory(dir string)
```

### Local Repository

The Local Repository package is used to parse and manage Maven local repository.

```go
package local_repository

// GetLocalRepositoryPath gets the local repository path
func GetLocalRepositoryPath() (string, error)

// ParseArtifactPath parses the artifact path
func ParseArtifactPath(groupId, artifactId, version string) string
```

### Installer

The Installer package provides automatic Maven installation functionality.

```go
package installer

// InstallMaven automatically installs the specified version of Maven
func InstallMaven(version string) error

// GetInstalledMavenVersion gets the installed Maven version
func GetInstalledMavenVersion() (string, error)
```

## Usage Examples

### Finding JAR Files

```go
jarPath, err := finder.FindJar("org.springframework", "spring-core", "5.3.21")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Spring Core JAR: %s\n", jarPath)
```

### Executing Maven Commands

```go
cmd := command.NewMavenCommand()
cmd.SetWorkingDirectory("/path/to/project")
err := cmd.Execute("clean", "install")
if err != nil {
    log.Fatal(err)
}
```

### Getting Local Repository Path

```go
repoPath, err := local_repository.GetLocalRepositoryPath()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Local repository path: %s\n", repoPath)
```

### Installing Maven

```go
err := installer.InstallMaven("3.8.6")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Maven installed successfully")
```