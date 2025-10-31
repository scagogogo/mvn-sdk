# Go Maven SDK

**Language**: [English](README.md) | [中文](README.zh.md)

## 1. What is this?

A Go SDK for conveniently operating `mvn` in Go. It detects and uses the locally installed `mvn` to execute commands.

## 2. Installation

```bash
go get -u github.com/scagogogo/mvn-sdk
```

## 3. Examples

Find locally installed mvn:

```go
package main

import (
	"fmt"
	"github.com/scagogogo/mvn-sdk/pkg/finder"
)

func main() {
	maven, err := finder.FindMaven()
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Maven executable path: %s", maven))
	// Output:
	// Maven executable path: mvn
}
```

Local repository directory:

```go
package main

import (
	"fmt"
	"github.com/scagogogo/mvn-sdk/pkg/finder"
	"github.com/scagogogo/mvn-sdk/pkg/local_repository"
)

func main() {
	maven, err := finder.FindMaven()
	if err != nil {
		panic(err)
	}
	directory := local_repository.ParseLocalRepositoryDirectory(maven)
	fmt.Println(fmt.Sprintf("Local repository location: %s", directory))
	// Output:
	// Local repository location: C:\Users\5950X\.m2\repository
}
```

Find JAR file location:

```go
package main

import (
	"fmt"
	"github.com/scagogogo/mvn-sdk/pkg/finder"
	"github.com/scagogogo/mvn-sdk/pkg/local_repository"
)

func main() {
	maven, err := finder.FindMaven()
	if err != nil {
		panic(err)
	}
	directory := local_repository.ParseLocalRepositoryDirectory(maven)
	jar, err := local_repository.FindJar(directory, "com.alibaba", "fastjson", "2.0.2")
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("JAR file location: %s", jar))
	// Output:
	// JAR file location: C:\Users\5950X\.m2\repository\com\alibaba\fastjson\2.0.2\fastjson-2.0.2.jar
}
```

Download JAR files. Note that if you have configured certain mirror repositories, the JAR files in the mirror repositories may not be complete:

```go
package main

import (
	"fmt"
	"github.com/scagogogo/mvn-sdk/pkg/command"
	"github.com/scagogogo/mvn-sdk/pkg/finder"
)

func main() {
	maven, err := finder.FindMaven()
	if err != nil {
		panic(err)
	}
	stdout, err := command.DependencyGet(maven, "com.alibaba", "fastjson", "2.0.2")
	if err != nil {
		panic(err)
	}
	fmt.Println(stdout)
	// Output:
	// [INFO] Scanning for projects...
	//[INFO]
	//[INFO] ------------------< org.apache.maven:standalone-pom >-------------------
	//[INFO] Building Maven Stub Project (No POM) 1
	//[INFO] --------------------------------[ pom ]---------------------------------
	//[INFO]
	//[INFO] --- maven-dependency-plugin:2.8:get (default-cli) @ standalone-pom ---
	//[INFO] Resolving com.alibaba:fastjson:jar:2.0.2 with transitive dependencies
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/alibaba/fastjson/1.2.80/fastjson-1.2.80.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/alibaba/fastjson/1.2.80/fastjson-1.2.80.pom (10 kB at 24 kB/s)
	// ...[truncated]
}
```

Execute arbitrary commands:

```go
package main

import (
	"fmt"
	"github.com/scagogogo/mvn-sdk/pkg/command"
	"github.com/scagogogo/mvn-sdk/pkg/finder"
)

func main() {
	maven, err := finder.FindMaven()
	if err != nil {
		panic(err)
	}
	stdout, err := command.ExecForStdout(maven, "-help")
	if err != nil {
		panic(err)
	}
	fmt.Println(stdout)
	// Output:
	// usage: mvn [options] [<goal(s)>] [<phase(s)>]
	//
	//Options:
	// -am,--also-make                        If project list is specified, also
	//                                        build projects required by the
	//                                        list
	// -amd,--also-make-dependents            If project list is specified, also
	//                                        build projects that depend on
	//                                        projects on the list
	// -B,--batch-mode                        Run in non-interactive (batch)
	//                                        mode (disables output color)
	// -b,--builder <arg>                     The id of the build strategy to
	//                                        use
	// -C,--strict-checksums                  Fail the build if checksums don't
	//                                        match
	// -c,--lax-checksums                     Warn if chec...[truncated]
}