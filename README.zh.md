# Go Maven SDK

**语言**: [English](README.md) | [中文](README.zh.md)

# 一、这是什么？

用于方便的在Go中操作`mvn`，会检测使用本地已经安装的`mvn`来执行命令。

# 二、安装依赖

```bash
go get -u github.com/scagogogo/mvn-sdk
```

# 三、文档

完整的API文档可在以下地址访问：https://scagogogo.github.io/mvn-sdk/

# 四、Example

寻找本地安装的mvn： 

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
	fmt.Println(fmt.Sprintf("Maven可执行路径： %s", maven))
	// Output:
	// Maven可执行路径： mvn

}
```

本地仓库目录：

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
	fmt.Println(fmt.Sprintf("本地仓库位置： %s", directory))
	// Output:
	// 本地仓库位置： C:\Users\5950X\.m2\repository

}
```

寻找jar包位置：

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
	fmt.Println(fmt.Sprintf("Jar包位置：%s", jar))
	// Output:
	// Jar包位置：C:\Users\5950X\.m2\repository\com\alibaba\fastjson\2.0.2\fastjson-2.0.2.jar

}
```

下载Jar包，需要注意如果配置了某些镜像仓库可能镜像仓库中的jar包不全：

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
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/retrofit2/retrofit/2.9.0/retrofit-2.9.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/retrofit2/retrofit/2.9.0/retrofit-2.9.0.pom (2.6 kB at 15 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okhttp3/okhttp/3.14.9/okhttp-3.14.9.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okhttp3/okhttp/3.14.9/okhttp-3.14.9.pom (2.5 kB at 16 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okhttp3/parent/3.14.9/parent-3.14.9.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okhttp3/parent/3.14.9/parent-3.14.9.pom (21 kB at 153 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okio/okio/1.17.2/okio-1.17.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okio/okio/1.17.2/okio-1.17.2.pom (2.0 kB at 14 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okio/okio-parent/1.17.2/okio-parent-1.17.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okio/okio-parent/1.17.2/okio-parent-1.17.2.pom (4.9 kB at 28 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/joda-time/joda-time/2.10.14/joda-time-2.10.14.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/joda-time/joda-time/2.10.14/joda-time-2.10.14.pom (37 kB at 260 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javamoney/moneta/moneta-core/1.4.2/moneta-core-1.4.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javamoney/moneta/moneta-core/1.4.2/moneta-core-1.4.2.pom (19 kB at 130 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javamoney/moneta-parent/1.4.2/moneta-parent-1.4.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javamoney/moneta-parent/1.4.2/moneta-parent-1.4.2.pom (20 kB at 127 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javamoney/javamoney-parent/1.3/javamoney-parent-1.3.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javamoney/javamoney-parent/1.3/javamoney-parent-1.3.pom (22 kB at 150 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/javax/money/money-api/1.1/money-api-1.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/javax/money/money-api/1.1/money-api-1.1.pom (32 kB at 157 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/airlift/slice/0.41/slice-0.41.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/airlift/slice/0.41/slice-0.41.pom (3.4 kB at 19 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-messaging/5.3.19/spring-messaging-5.3.19.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-messaging/5.3.19/spring-messaging-5.3.19.pom (2.2 kB at 10 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-websocket/5.3.19/spring-websocket-5.3.19.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-websocket/5.3.19/spring-websocket-5.3.19.pom (2.4 kB at 17 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-redis/2.6.4/spring-data-redis-2.6.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-redis/2.6.4/spring-data-redis-2.6.4.pom (8.0 kB at 58 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/build/spring-data-parent/2.6.4/spring-data-parent-2.6.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/build/spring-data-parent/2.6.4/spring-data-parent-2.6.4.pom (39 kB at 275 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/build/spring-data-build/2.6.4/spring-data-build-2.6.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/build/spring-data-build/2.6.4/spring-data-build-2.6.4.pom (7.3 kB at 54 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/projectreactor/reactor-bom/2020.0.18/reactor-bom-2020.0.18.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/projectreactor/reactor-bom/2020.0.18/reactor-bom-2020.0.18.pom (4.6 kB at 35 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-bom/1.5.32/kotlin-bom-1.5.32.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-bom/1.5.32/kotlin-bom-1.5.32.pom (9.3 kB at 72 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jackson-bom/2.13.2.1/jackson-bom-2.13.2.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jackson-bom/2.13.2.1/jackson-bom-2.13.2.1.pom (17 kB at 137 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-keyvalue/2.6.4/spring-data-keyvalue-2.6.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-keyvalue/2.6.4/spring-data-keyvalue-2.6.4.pom (2.8 kB at 22 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-commons/2.6.4/spring-data-commons-2.6.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-commons/2.6.4/spring-data-commons-2.6.4.pom (10 kB at 71 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-tx/5.3.19/spring-tx-5.3.19.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-tx/5.3.19/spring-tx-5.3.19.pom (2.2 kB at 17 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-oxm/5.3.19/spring-oxm-5.3.19.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-oxm/5.3.19/spring-oxm-5.3.19.pom (2.2 kB at 17 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-context-support/5.3.19/spring-context-support-5.3.19.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-context-support/5.3.19/spring-context-support-5.3.19.pom (2.4 kB at 18 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-databind/2.13.2/jackson-databind-2.13.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-databind/2.13.2/jackson-databind-2.13.2.pom (16 kB at 98 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-module-afterburner/2.13.2/jackson-module-afterburner-2.13.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-module-afterburner/2.13.2/jackson-module-afterburner-2.13.2.pom (3.4 kB at 24 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-module-kotlin/2.13.2/jackson-module-kotlin-2.13.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-module-kotlin/2.13.2/jackson-module-kotlin-2.13.2.pom (9.4 kB at 69 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-reflect/1.5.30/kotlin-reflect-1.5.30.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-reflect/1.5.30/kotlin-reflect-1.5.30.pom (1.4 kB at 8.2 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib/1.5.30/kotlin-stdlib-1.5.30.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib/1.5.30/kotlin-stdlib-1.5.30.pom (1.6 kB at 12 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/annotations/13.0/annotations-13.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/annotations/13.0/annotations-13.0.pom (4.9 kB at 35 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib-common/1.5.30/kotlin-stdlib-common-1.5.30.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib-common/1.5.30/kotlin-stdlib-common-1.5.30.pom (1.2 kB at 8.7 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/code/gson/gson/2.9.0/gson-2.9.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/code/gson/gson/2.9.0/gson-2.9.0.pom (8.1 kB at 57 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/code/gson/gson-parent/2.9.0/gson-parent-2.9.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/code/gson/gson-parent/2.9.0/gson-parent-2.9.0.pom (4.5 kB at 32 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/junit/junit/4.10/junit-4.10.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/junit/junit/4.10/junit-4.10.pom (2.3 kB at 16 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/sf/json-lib/json-lib/2.4/json-lib-2.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/sf/json-lib/json-lib/2.4/json-lib-2.4.pom (13 kB at 102 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/sf/ezmorph/ezmorph/1.0.6/ezmorph-1.0.6.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/sf/ezmorph/ezmorph/1.0.6/ezmorph-1.0.6.pom (6.8 kB at 50 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/commons-lang/commons-lang/2.3/commons-lang-2.3.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/commons-lang/commons-lang/2.3/commons-lang-2.3.pom (11 kB at 88 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/minidev/json-smart/2.4.8/json-smart-2.4.8.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/minidev/json-smart/2.4.8/json-smart-2.4.8.pom (8.3 kB at 61 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/minidev/accessors-smart/2.4.8/accessors-smart-2.4.8.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/minidev/accessors-smart/2.4.8/accessors-smart-2.4.8.pom (10 kB at 78 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/slf4j/slf4j-api/1.7.33/slf4j-api-1.7.33.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/slf4j/slf4j-api/1.7.33/slf4j-api-1.7.33.pom (3.8 kB at 21 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/slf4j/slf4j-parent/1.7.33/slf4j-parent-1.7.33.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/slf4j/slf4j-parent/1.7.33/slf4j-parent-1.7.33.pom (14 kB at 108 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/protobuf/protobuf-java/3.20.1/protobuf-java-3.20.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/protobuf/protobuf-java/3.20.1/protobuf-java-3.20.1.pom (5.4 kB at 41 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/protobuf/protobuf-parent/3.20.1/protobuf-parent-3.20.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/protobuf/protobuf-parent/3.20.1/protobuf-parent-3.20.1.pom (9.0 kB at 67 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/protobuf/protobuf-java-util/3.20.1/protobuf-java-util-3.20.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/protobuf/protobuf-java-util/3.20.1/protobuf-java-util-3.20.1.pom (5.6 kB at 42 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib/1.4.32/kotlin-stdlib-1.4.32.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib/1.4.32/kotlin-stdlib-1.4.32.pom (1.6 kB at 12 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib-common/1.4.32/kotlin-stdlib-common-1.4.32.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib-common/1.4.32/kotlin-stdlib-common-1.4.32.pom (1.2 kB at 8.9 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-reflect/1.4.32/kotlin-reflect-1.4.32.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-reflect/1.4.32/kotlin-reflect-1.4.32.pom (1.4 kB at 11 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/clojure/clojure/1.5.1/clojure-1.5.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/clojure/clojure/1.5.1/clojure-1.5.1.pom (6.2 kB at 31 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/codehaus/groovy/groovy/2.1.5/groovy-2.1.5.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/codehaus/groovy/groovy/2.1.5/groovy-2.1.5.pom (16 kB at 86 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps-sdk-udf/0.38.4-public/odps-sdk-udf-0.38.4-public.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps-sdk-udf/0.38.4-public/odps-sdk-udf-0.38.4-public.pom (1.2 kB at 4.3 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps-sdk/0.38.4-public/odps-sdk-0.38.4-public.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps-sdk/0.38.4-public/odps-sdk-0.38.4-public.pom (1.2 kB at 6.5 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps/0.38.4-public/odps-0.38.4-public.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps/0.38.4-public/odps-0.38.4-public.pom (20 kB at 106 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps-sdk-commons/0.38.4-public/odps-sdk-commons-0.38.4-public.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps-sdk-commons/0.38.4-public/odps-sdk-commons-0.38.4-public.pom (2.9 kB at 15 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/aspectj/aspectjrt/1.8.9/aspectjrt-1.8.9.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/aspectj/aspectjrt/1.8.9/aspectjrt-1.8.9.pom (1.0 kB at 6.9 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/hibernate-core/5.2.10.Final/hibernate-core-5.2.10.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/hibernate-core/5.2.10.Final/hibernate-core-5.2.10.Final.pom (3.5 kB at 17 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/logging/jboss-logging/3.3.0.Final/jboss-logging-3.3.0.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/logging/jboss-logging/3.3.0.Final/jboss-logging-3.3.0.Final.pom (5.9 kB at 45 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/javax/persistence/hibernate-jpa-2.1-api/1.0.0.Final/hibernate-jpa-2.1-api-1.0.0.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/javax/persistence/hibernate-jpa-2.1-api/1.0.0.Final/hibernate-jpa-2.1-api-1.0.0.Final.pom (2.2 kB at 16 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.20.0-GA/javassist-3.20.0-GA.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.20.0-GA/javassist-3.20.0-GA.pom (9.8 kB at 76 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/spec/javax/transaction/jboss-transaction-api_1.2_spec/1.0.1.Final/jboss-transaction-api_1.2_spec-1.0.1.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/spec/javax/transaction/jboss-transaction-api_1.2_spec/1.0.1.Final/jboss-transaction-api_1.2_spec-1.0.1.Final.pom (5.0 kB at 30 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/jboss-parent/20/jboss-parent-20.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/jboss-parent/20/jboss-parent-20.pom (34 kB at 241 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/jandex/2.0.3.Final/jandex-2.0.3.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/jandex/2.0.3.Final/jandex-2.0.3.Final.pom (5.7 kB at 38 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/jboss-parent/12/jboss-parent-12.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/jboss-parent/12/jboss-parent-12.pom (32 kB at 248 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/classmate/1.3.0/classmate-1.3.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/classmate/1.3.0/classmate-1.3.0.pom (5.7 kB at 33 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/dom4j/dom4j/1.6.1/dom4j-1.6.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/dom4j/dom4j/1.6.1/dom4j-1.6.1.pom (6.8 kB at 45 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/common/hibernate-commons-annotations/5.0.1.Final/hibernate-commons-annotations-5.0.1.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/common/hibernate-commons-annotations/5.0.1.Final/hibernate-commons-annotations-5.0.1.Final.pom (1.9 kB at 13 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/gitlab4j/gitlab4j-api/5.0.1/gitlab4j-api-5.0.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/gitlab4j/gitlab4j-api/5.0.1/gitlab4j-api-5.0.1.pom (13 kB at 82 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/inject/jersey-hk2/2.35/jersey-hk2-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/inject/jersey-hk2/2.35/jersey-hk2-2.35.pom (4.7 kB at 38 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/inject/project/2.35/project-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/inject/project/2.35/project-2.35.pom (1.5 kB at 12 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/project/2.35/project-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/project/2.35/project-2.35.pom (99 kB at 511 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-common/2.35/jersey-common-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-common/2.35/jersey-common-2.35.pom (35 kB at 268 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/jakarta/ws/rs/jakarta.ws.rs-api/2.1.6/jakarta.ws.rs-api-2.1.6.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/jakarta/ws/rs/jakarta.ws.rs-api/2.1.6/jakarta.ws.rs-api-2.1.6.pom (35 kB at 269 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/jakarta.inject/2.6.1/jakarta.inject-2.6.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/jakarta.inject/2.6.1/jakarta.inject-2.6.1.pom (5.2 kB at 41 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/2.6.1/external-2.6.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/2.6.1/external-2.6.1.pom (1.5 kB at 12 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-parent/2.6.1/hk2-parent-2.6.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-parent/2.6.1/hk2-parent-2.6.1.pom (42 kB at 323 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/osgi-resource-locator/1.0.3/osgi-resource-locator-1.0.3.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/osgi-resource-locator/1.0.3/osgi-resource-locator-1.0.3.pom (7.4 kB at 42 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-locator/2.6.1/hk2-locator-2.6.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-locator/2.6.1/hk2-locator-2.6.1.pom (6.5 kB at 45 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/aopalliance-repackaged/2.6.1/aopalliance-repackaged-2.6.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/aopalliance-repackaged/2.6.1/aopalliance-repackaged-2.6.1.pom (5.5 kB at 42 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-api/2.6.1/hk2-api-2.6.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-api/2.6.1/hk2-api-2.6.1.pom (3.4 kB at 26 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-utils/2.6.1/hk2-utils-2.6.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-utils/2.6.1/hk2-utils-2.6.1.pom (4.9 kB at 37 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.25.0-GA/javassist-3.25.0-GA.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.25.0-GA/javassist-3.25.0-GA.pom (11 kB at 68 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-client/2.35/jersey-client-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-client/2.35/jersey-client-2.35.pom (6.8 kB at 53 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/connectors/jersey-apache-connector/2.35/jersey-apache-connector-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/connectors/jersey-apache-connector/2.35/jersey-apache-connector-2.35.pom (3.2 kB at 26 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/connectors/project/2.35/project-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/connectors/project/2.35/project-2.35.pom (3.1 kB at 25 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-multipart/2.35/jersey-media-multipart-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-multipart/2.35/jersey-media-multipart-2.35.pom (3.9 kB at 30 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/project/2.35/project-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/project/2.35/project-2.35.pom (2.0 kB at 15 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jvnet/mimepull/mimepull/1.9.13/mimepull-1.9.13.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jvnet/mimepull/mimepull/1.9.13/mimepull-1.9.13.pom (16 kB at 111 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-json-jackson/2.35/jersey-media-json-jackson-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-json-jackson/2.35/jersey-media-json-jackson-2.35.pom (5.9 kB at 45 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/ext/jersey-entity-filtering/2.35/jersey-entity-filtering-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/ext/jersey-entity-filtering/2.35/jersey-entity-filtering-2.35.pom (3.5 kB at 28 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/ext/project/2.35/project-2.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/ext/project/2.35/project-2.35.pom (2.8 kB at 21 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-annotations/2.12.2/jackson-annotations-2.12.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-annotations/2.12.2/jackson-annotations-2.12.2.pom (6.0 kB at 46 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-databind/2.12.2/jackson-databind-2.12.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-databind/2.12.2/jackson-databind-2.12.2.pom (15 kB at 117 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jackson-base/2.12.2/jackson-base-2.12.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jackson-base/2.12.2/jackson-base-2.12.2.pom (9.3 kB at 67 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jackson-bom/2.12.2/jackson-bom-2.12.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jackson-bom/2.12.2/jackson-bom-2.12.2.pom (17 kB at 132 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-module-jaxb-annotations/2.12.2/jackson-module-jaxb-annotations-2.12.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-module-jaxb-annotations/2.12.2/jackson-module-jaxb-annotations-2.12.2.pom (5.3 kB at 41 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-modules-base/2.12.2/jackson-modules-base-2.12.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-modules-base/2.12.2/jackson-modules-base-2.12.2.pom (3.5 kB at 26 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/jakarta/servlet/jakarta.servlet-api/4.0.4/jakarta.servlet-api-4.0.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/jakarta/servlet/jakarta.servlet-api/4.0.4/jakarta.servlet-api-4.0.4.pom (17 kB at 130 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/uk/org/webcompere/system-stubs-jupiter/1.2.0/system-stubs-jupiter-1.2.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/uk/org/webcompere/system-stubs-jupiter/1.2.0/system-stubs-jupiter-1.2.0.pom (3.9 kB at 30 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/uk/org/webcompere/system-stubs-parent/1.2.0/system-stubs-parent-1.2.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/uk/org/webcompere/system-stubs-parent/1.2.0/system-stubs-parent-1.2.0.pom (11 kB at 73 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/uk/org/webcompere/system-stubs-core/1.2.0/system-stubs-core-1.2.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/uk/org/webcompere/system-stubs-core/1.2.0/system-stubs-core-1.2.0.pom (2.6 kB at 19 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/chinamobile/cmos/sms-core/2.1.12.5/sms-core-2.1.12.5.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/chinamobile/cmos/sms-core/2.1.12.5/sms-core-2.1.12.5.pom (8.6 kB at 28 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/de/ruedigermoeller/fst/2.48-jdk-6/fst-2.48-jdk-6.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/de/ruedigermoeller/fst/2.48-jdk-6/fst-2.48-jdk-6.pom (7.5 kB at 38 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-core/2.8.6/jackson-core-2.8.6.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-core/2.8.6/jackson-core-2.8.6.pom (5.4 kB at 44 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.19.0-GA/javassist-3.19.0-GA.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.19.0-GA/javassist-3.19.0-GA.pom (9.6 kB at 72 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/objenesis/objenesis/2.4/objenesis-2.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/objenesis/objenesis/2.4/objenesis-2.4.pom (2.4 kB at 16 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/objenesis/objenesis-parent/2.4/objenesis-parent-2.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/objenesis/objenesis-parent/2.4/objenesis-parent-2.4.pom (17 kB at 101 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/cedarsoftware/java-util/1.9.0/java-util-1.9.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/cedarsoftware/java-util/1.9.0/java-util-1.9.0.pom (4.9 kB at 33 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/guava/guava/16.0/guava-16.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/guava/guava/16.0/guava-16.0.pom (6.1 kB at 44 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/guava/guava-parent/16.0/guava-parent-16.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/guava/guava-parent/16.0/guava-parent-16.0.pom (7.3 kB at 60 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-transport/4.1.51.Final/netty-transport-4.1.51.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-transport/4.1.51.Final/netty-transport-4.1.51.Final.pom (1.9 kB at 16 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-parent/4.1.51.Final/netty-parent-4.1.51.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-parent/4.1.51.Final/netty-parent-4.1.51.Final.pom (58 kB at 413 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-common/4.1.51.Final/netty-common-4.1.51.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-common/4.1.51.Final/netty-common-4.1.51.Final.pom (10 kB at 82 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-buffer/4.1.51.Final/netty-buffer-4.1.51.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-buffer/4.1.51.Final/netty-buffer-4.1.51.Final.pom (1.6 kB at 13 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-resolver/4.1.51.Final/netty-resolver-4.1.51.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-resolver/4.1.51.Final/netty-resolver-4.1.51.Final.pom (1.6 kB at 12 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec/4.1.51.Final/netty-codec-4.1.51.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec/4.1.51.Final/netty-codec-4.1.51.Final.pom (3.6 kB at 29 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-handler-proxy/4.1.51.Final/netty-handler-proxy-4.1.51.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-handler-proxy/4.1.51.Final/netty-handler-proxy-4.1.51.Final.pom (2.8 kB at 21 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec-socks/4.1.51.Final/netty-codec-socks-4.1.51.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec-socks/4.1.51.Final/netty-codec-socks-4.1.51.Final.pom (2.0 kB at 15 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec-http/4.1.51.Final/netty-codec-http-4.1.51.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec-http/4.1.51.Final/netty-codec-http-4.1.51.Final.pom (2.4 kB at 17 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-handler/4.1.51.Final/netty-handler-4.1.51.Final.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-handler/4.1.51.Final/netty-handler-4.1.51.Final.pom (3.6 kB at 28 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/sleepycat/je/18.3.12/je-18.3.12.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/sleepycat/je/18.3.12/je-18.3.12.pom (1.3 kB at 5.7 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/commons-io/commons-io/2.7/commons-io-2.7.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/commons-io/commons-io/2.7/commons-io-2.7.pom (16 kB at 122 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/apache/commons/commons-parent/50/commons-parent-50.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/apache/commons/commons-parent/50/commons-parent-50.pom (76 kB at 573 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/cglib/cglib-nodep/3.3.0/cglib-nodep-3.3.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/cglib/cglib-nodep/3.3.0/cglib-nodep-3.3.0.pom (4.3 kB at 31 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/cglib/cglib-parent/3.3.0/cglib-parent-3.3.0.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/cglib/cglib-parent/3.3.0/cglib-parent-3.3.0.pom (10 kB at 84 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.28.0-GA/javassist-3.28.0-GA.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.28.0-GA/javassist-3.28.0-GA.pom (11 kB at 86 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-server/9.4.17.v20190418/jetty-server-9.4.17.v20190418.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-server/9.4.17.v20190418/jetty-server-9.4.17.v20190418.pom (2.6 kB at 12 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-project/9.4.17.v20190418/jetty-project-9.4.17.v20190418.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-project/9.4.17.v20190418/jetty-project-9.4.17.v20190418.pom (70 kB at 332 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-http/9.4.17.v20190418/jetty-http-9.4.17.v20190418.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-http/9.4.17.v20190418/jetty-http-9.4.17.v20190418.pom (4.3 kB at 19 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-util/9.4.17.v20190418/jetty-util-9.4.17.v20190418.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-util/9.4.17.v20190418/jetty-util-9.4.17.v20190418.pom (3.8 kB at 18 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-io/9.4.17.v20190418/jetty-io-9.4.17.v20190418.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-io/9.4.17.v20190418/jetty-io-9.4.17.v20190418.pom (1.3 kB at 6.6 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-webapp/9.4.17.v20190418/jetty-webapp-9.4.17.v20190418.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-webapp/9.4.17.v20190418/jetty-webapp-9.4.17.v20190418.pom (3.0 kB at 15 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-xml/9.4.17.v20190418/jetty-xml-9.4.17.v20190418.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-xml/9.4.17.v20190418/jetty-xml-9.4.17.v20190418.pom (1.9 kB at 9.4 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-servlet/9.4.17.v20190418/jetty-servlet-9.4.17.v20190418.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-servlet/9.4.17.v20190418/jetty-servlet-9.4.17.v20190418.pom (2.1 kB at 8.4 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-security/9.4.17.v20190418/jetty-security-9.4.17.v20190418.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-security/9.4.17.v20190418/jetty-security-9.4.17.v20190418.pom (1.6 kB at 6.5 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-common/2.23.2/jersey-common-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-common/2.23.2/jersey-common-2.23.2.pom (10 kB at 82 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/project/2.23.2/project-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/project/2.23.2/project-2.23.2.pom (87 kB at 638 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-bom/2.5.0-b05/hk2-bom-2.5.0-b05.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-bom/2.5.0-b05/hk2-bom-2.5.0-b05.pom (20 kB at 128 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/bundles/repackaged/jersey-guava/2.23.2/jersey-guava-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/bundles/repackaged/jersey-guava/2.23.2/jersey-guava-2.23.2.pom (13 kB at 100 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/bundles/repackaged/project/2.23.2/project-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/bundles/repackaged/project/2.23.2/project-2.23.2.pom (2.8 kB at 20 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/bundles/project/2.23.2/project-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/bundles/project/2.23.2/project-2.23.2.pom (3.1 kB at 24 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-api/2.5.0-b05/hk2-api-2.5.0-b05.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-api/2.5.0-b05/hk2-api-2.5.0-b05.pom (4.6 kB at 34 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-parent/2.5.0-b05/hk2-parent-2.5.0-b05.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-parent/2.5.0-b05/hk2-parent-2.5.0-b05.pom (49 kB at 289 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-utils/2.5.0-b05/hk2-utils-2.5.0-b05.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-utils/2.5.0-b05/hk2-utils-2.5.0-b05.pom (5.3 kB at 41 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/aopalliance-repackaged/2.5.0-b05/aopalliance-repackaged-2.5.0-b05.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/aopalliance-repackaged/2.5.0-b05/aopalliance-repackaged-2.5.0-b05.pom (6.8 kB at 46 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/2.5.0-b05/external-2.5.0-b05.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/2.5.0-b05/external-2.5.0-b05.pom (2.9 kB at 25 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/javax.inject/2.5.0-b05/javax.inject-2.5.0-b05.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/javax.inject/2.5.0-b05/javax.inject-2.5.0-b05.pom (6.5 kB at 42 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-locator/2.5.0-b05/hk2-locator-2.5.0-b05.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-locator/2.5.0-b05/hk2-locator-2.5.0-b05.pom (7.5 kB at 57 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-servlet/2.23.2/jersey-container-servlet-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-servlet/2.23.2/jersey-container-servlet-2.23.2.pom (4.6 kB at 23 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/project/2.23.2/project-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/project/2.23.2/project-2.23.2.pom (3.8 kB at 19 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-servlet-core/2.23.2/jersey-container-servlet-core-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-servlet-core/2.23.2/jersey-container-servlet-core-2.23.2.pom (4.7 kB at 21 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-server/2.23.2/jersey-server-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-server/2.23.2/jersey-server-2.23.2.pom (12 kB at 43 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-client/2.23.2/jersey-client-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-client/2.23.2/jersey-client-2.23.2.pom (7.6 kB at 59 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-jaxb/2.23.2/jersey-media-jaxb-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-jaxb/2.23.2/jersey-media-jaxb-2.23.2.pom (7.1 kB at 36 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/project/2.23.2/project-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/project/2.23.2/project-2.23.2.pom (3.0 kB at 15 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/providers/jersey-test-framework-provider-jdk-http/2.23.2/jersey-test-framework-provider-jdk-http-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/providers/jersey-test-framework-provider-jdk-http/2.23.2/jersey-test-framework-provider-jdk-http-2.23.2.pom (3.6 kB at 18 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/providers/project/2.23.2/project-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/providers/project/2.23.2/project-2.23.2.pom (3.0 kB at 16 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/project/2.23.2/project-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/project/2.23.2/project-2.23.2.pom (3.2 kB at 16 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/jersey-test-framework-core/2.23.2/jersey-test-framework-core-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/jersey-test-framework-core/2.23.2/jersey-test-framework-core-2.23.2.pom (4.4 kB at 21 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/javax/servlet/javax.servlet-api/3.0.1/javax.servlet-api-3.0.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/javax/servlet/javax.servlet-api/3.0.1/javax.servlet-api-3.0.1.pom (13 kB at 105 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/ow2/asm/asm-debug-all/5.0.4/asm-debug-all-5.0.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/ow2/asm/asm-debug-all/5.0.4/asm-debug-all-5.0.4.pom (2.0 kB at 9.7 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/ow2/asm/asm-parent/5.0.4/asm-parent-5.0.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/ow2/asm/asm-parent/5.0.4/asm-parent-5.0.4.pom (5.5 kB at 50 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-jdk-http/2.23.2/jersey-container-jdk-http-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-jdk-http/2.23.2/jersey-container-jdk-http-2.23.2.pom (5.5 kB at 27 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-json-jackson/2.23.2/jersey-media-json-jackson-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-json-jackson/2.23.2/jersey-media-json-jackson-2.23.2.pom (4.8 kB at 23 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/ext/jersey-entity-filtering/2.23.2/jersey-entity-filtering-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/ext/jersey-entity-filtering/2.23.2/jersey-entity-filtering-2.23.2.pom (4.6 kB at 21 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/ext/project/2.23.2/project-2.23.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/ext/project/2.23.2/project-2.23.2.pom (3.8 kB at 19 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jaxrs/jackson-jaxrs-base/2.5.4/jackson-jaxrs-base-2.5.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jaxrs/jackson-jaxrs-base/2.5.4/jackson-jaxrs-base-2.5.4.pom (1.9 kB at 15 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jaxrs/jackson-jaxrs-providers/2.5.4/jackson-jaxrs-providers-2.5.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jaxrs/jackson-jaxrs-providers/2.5.4/jackson-jaxrs-providers-2.5.4.pom (3.9 kB at 34 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jackson-parent/2.5.1/jackson-parent-2.5.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jackson-parent/2.5.1/jackson-parent-2.5.1.pom (7.8 kB at 67 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/oss-parent/19/oss-parent-19.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/oss-parent/19/oss-parent-19.pom (19 kB at 159 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jaxrs/jackson-jaxrs-json-provider/2.5.4/jackson-jaxrs-json-provider-2.5.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/jaxrs/jackson-jaxrs-json-provider/2.5.4/jackson-jaxrs-json-provider-2.5.4.pom (3.4 kB at 29 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-annotations/2.5.4/jackson-annotations-2.5.4.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-annotations/2.5.4/jackson-annotations-2.5.4.pom (1.2 kB at 9.9 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-test/5.3.19/spring-test-5.3.19.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-test/5.3.19/spring-test-5.3.19.pom (2.1 kB at 16 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-web/5.6.3/spring-security-web-5.6.3.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-web/5.6.3/spring-security-web-5.6.3.pom (3.2 kB at 26 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-core/5.6.3/spring-security-core-5.6.3.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-core/5.6.3/spring-security-core-5.6.3.pom (3.0 kB at 23 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-crypto/5.6.3/spring-security-crypto-5.6.3.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-crypto/5.6.3/spring-security-crypto-5.6.3.pom (1.9 kB at 9.2 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/junit/vintage/junit-vintage-engine/5.8.2/junit-vintage-engine-5.8.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/junit/vintage/junit-vintage-engine/5.8.2/junit-vintage-engine-5.8.2.pom (0 B at 0 B/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/mockito/mockito-all/1.10.19/mockito-all-1.10.19.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/mockito/mockito-all/1.10.19/mockito-all-1.10.19.pom (930 B at 7.8 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-module-junit4/2.0.9/powermock-module-junit4-2.0.9.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-module-junit4/2.0.9/powermock-module-junit4-2.0.9.pom (8.0 kB at 60 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-module-junit4-common/2.0.9/powermock-module-junit4-common-2.0.9.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-module-junit4-common/2.0.9/powermock-module-junit4-common-2.0.9.pom (8.2 kB at 65 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-reflect/2.0.9/powermock-reflect-2.0.9.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-reflect/2.0.9/powermock-reflect-2.0.9.pom (7.9 kB at 59 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/objenesis/objenesis/3.0.1/objenesis-3.0.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/objenesis/objenesis/3.0.1/objenesis-3.0.1.pom (3.5 kB at 28 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/objenesis/objenesis-parent/3.0.1/objenesis-parent-3.0.1.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/objenesis/objenesis-parent/3.0.1/objenesis-parent-3.0.1.pom (17 kB at 131 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/bytebuddy/byte-buddy/1.10.14/byte-buddy-1.10.14.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/bytebuddy/byte-buddy/1.10.14/byte-buddy-1.10.14.pom (11 kB at 97 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/bytebuddy/byte-buddy-parent/1.10.14/byte-buddy-parent-1.10.14.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/bytebuddy/byte-buddy-parent/1.10.14/byte-buddy-parent-1.10.14.pom (41 kB at 329 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/bytebuddy/byte-buddy-agent/1.10.14/byte-buddy-agent-1.10.14.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/bytebuddy/byte-buddy-agent/1.10.14/byte-buddy-agent-1.10.14.pom (9.6 kB at 86 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-core/2.0.9/powermock-core-2.0.9.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-core/2.0.9/powermock-core-2.0.9.pom (8.1 kB at 60 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.27.0-GA/javassist-3.27.0-GA.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.27.0-GA/javassist-3.27.0-GA.pom (11 kB at 92 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/openjdk/jmh/jmh-core/1.35/jmh-core-1.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/openjdk/jmh/jmh-core/1.35/jmh-core-1.35.pom (9.6 kB at 74 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/openjdk/jmh/jmh-parent/1.35/jmh-parent-1.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/openjdk/jmh/jmh-parent/1.35/jmh-parent-1.35.pom (11 kB at 81 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/apache/commons/commons-math3/3.2/commons-math3-3.2.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/apache/commons/commons-math3/3.2/commons-math3-3.2.pom (17 kB at 144 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/openjdk/jmh/jmh-generator-annprocess/1.35/jmh-generator-annprocess-1.35.pom
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/openjdk/jmh/jmh-generator-annprocess/1.35/jmh-generator-annprocess-1.35.pom (3.8 kB at 25 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/joda-time/joda-time/2.10.14/joda-time-2.10.14.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okio/okio/1.17.2/okio-1.17.2.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/retrofit2/retrofit/2.9.0/retrofit-2.9.0.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javamoney/moneta/moneta-core/1.4.2/moneta-core-1.4.2.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okhttp3/okhttp/3.14.9/okhttp-3.14.9.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okio/okio/1.17.2/okio-1.17.2.jar (92 kB at 477 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/javax/money/money-api/1.1/money-api-1.1.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/joda-time/joda-time/2.10.14/joda-time-2.10.14.jar (644 kB at 2.4 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/airlift/slice/0.41/slice-0.41.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/retrofit2/retrofit/2.9.0/retrofit-2.9.0.jar (125 kB at 379 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-messaging/5.3.19/spring-messaging-5.3.19.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/squareup/okhttp3/okhttp/3.14.9/okhttp-3.14.9.jar (430 kB at 1.3 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-websocket/5.3.19/spring-websocket-5.3.19.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/javax/money/money-api/1.1/money-api-1.1.jar (85 kB at 422 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-redis/2.6.4/spring-data-redis-2.6.4.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javamoney/moneta/moneta-core/1.4.2/moneta-core-1.4.2.jar (233 kB at 524 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-keyvalue/2.6.4/spring-data-keyvalue-2.6.4.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/airlift/slice/0.41/slice-0.41.jar (71 kB at 343 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-commons/2.6.4/spring-data-commons-2.6.4.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-keyvalue/2.6.4/spring-data-keyvalue-2.6.4.jar (121 kB at 653 kB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-messaging/5.3.19/spring-messaging-5.3.19.jar (567 kB at 1.9 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-tx/5.3.19/spring-tx-5.3.19.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-oxm/5.3.19/spring-oxm-5.3.19.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-websocket/5.3.19/spring-websocket-5.3.19.jar (448 kB at 1.3 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-context-support/5.3.19/spring-context-support-5.3.19.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-redis/2.6.4/spring-data-redis-2.6.4.jar (2.1 MB at 6.8 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-databind/2.13.2/jackson-databind-2.13.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/data/spring-data-commons/2.6.4/spring-data-commons-2.6.4.jar (1.3 MB at 5.0 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-module-afterburner/2.13.2/jackson-module-afterburner-2.13.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-tx/5.3.19/spring-tx-5.3.19.jar (333 kB at 1.5 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-module-kotlin/2.13.2/jackson-module-kotlin-2.13.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-context-support/5.3.19/spring-context-support-5.3.19.jar (187 kB at 1.0 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-reflect/1.4.32/kotlin-reflect-1.4.32.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-oxm/5.3.19/spring-oxm-5.3.19.jar (66 kB at 280 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib/1.4.32/kotlin-stdlib-1.4.32.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-module-afterburner/2.13.2/jackson-module-afterburner-2.13.2.jar (215 kB at 1.1 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/annotations/13.0/annotations-13.0.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/core/jackson-databind/2.13.2/jackson-databind-2.13.2.jar (1.5 MB at 6.6 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib-common/1.4.32/kotlin-stdlib-common-1.4.32.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/jackson/module/jackson-module-kotlin/2.13.2/jackson-module-kotlin-2.13.2.jar (147 kB at 857 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/code/gson/gson/2.9.0/gson-2.9.0.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/annotations/13.0/annotations-13.0.jar (18 kB at 110 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/sf/json-lib/json-lib/2.4/json-lib-2.4-jdk15.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib-common/1.4.32/kotlin-stdlib-common-1.4.32.jar (193 kB at 975 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/commons-beanutils/commons-beanutils/1.8.0/commons-beanutils-1.8.0.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-stdlib/1.4.32/kotlin-stdlib-1.4.32.jar (1.5 MB at 4.1 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/sf/ezmorph/ezmorph/1.0.6/ezmorph-1.0.6.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/code/gson/gson/2.9.0/gson-2.9.0.jar (249 kB at 1.1 MB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jetbrains/kotlin/kotlin-reflect/1.4.32/kotlin-reflect-1.4.32.jar (3.0 MB at 7.7 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/minidev/json-smart/2.4.8/json-smart-2.4.8.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/minidev/accessors-smart/2.4.8/accessors-smart-2.4.8.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/sf/json-lib/json-lib/2.4/json-lib-2.4-jdk15.jar (159 kB at 936 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/protobuf/protobuf-java/3.20.1/protobuf-java-3.20.1.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/commons-beanutils/commons-beanutils/1.8.0/commons-beanutils-1.8.0.jar (231 kB at 1.4 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/protobuf/protobuf-java-util/3.20.1/protobuf-java-util-3.20.1.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/sf/ezmorph/ezmorph/1.0.6/ezmorph-1.0.6.jar (86 kB at 703 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/clojure/clojure/1.5.1/clojure-1.5.1.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/minidev/accessors-smart/2.4.8/accessors-smart-2.4.8.jar (30 kB at 180 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/codehaus/groovy/groovy/2.1.5/groovy-2.1.5.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/minidev/json-smart/2.4.8/json-smart-2.4.8.jar (120 kB at 543 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps-sdk-udf/0.38.4-public/odps-sdk-udf-0.38.4-public.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/protobuf/protobuf-java-util/3.20.1/protobuf-java-util-3.20.1.jar (74 kB at 396 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps-sdk-commons/0.38.4-public/odps-sdk-commons-0.38.4-public.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/google/protobuf/protobuf-java/3.20.1/protobuf-java-3.20.1.jar (1.7 MB at 6.7 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/aspectj/aspectjrt/1.8.9/aspectjrt-1.8.9.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/aspectj/aspectjrt/1.8.9/aspectjrt-1.8.9.jar (118 kB at 450 kB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps-sdk-udf/0.38.4-public/odps-sdk-udf-0.38.4-public.jar (36 kB at 114 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/hibernate-core/5.2.10.Final/hibernate-core-5.2.10.Final.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/logging/jboss-logging/3.3.0.Final/jboss-logging-3.3.0.Final.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/aliyun/odps/odps-sdk-commons/0.38.4-public/odps-sdk-commons-0.38.4-public.jar (196 kB at 569 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/javax/persistence/hibernate-jpa-2.1-api/1.0.0.Final/hibernate-jpa-2.1-api-1.0.0.Final.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/clojure/clojure/1.5.1/clojure-1.5.1.jar (3.6 MB at 6.9 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.28.0-GA/javassist-3.28.0-GA.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/logging/jboss-logging/3.3.0.Final/jboss-logging-3.3.0.Final.jar (67 kB at 373 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/spec/javax/transaction/jboss-transaction-api_1.2_spec/1.0.1.Final/jboss-transaction-api_1.2_spec-1.0.1.Final.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/codehaus/groovy/groovy/2.1.5/groovy-2.1.5.jar (3.4 MB at 5.8 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/jandex/2.0.3.Final/jandex-2.0.3.Final.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/javax/persistence/hibernate-jpa-2.1-api/1.0.0.Final/hibernate-jpa-2.1-api-1.0.0.Final.jar (113 kB at 476 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/classmate/1.3.0/classmate-1.3.0.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/javassist/javassist/3.28.0-GA/javassist-3.28.0-GA.jar (852 kB at 3.0 MB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/spec/javax/transaction/jboss-transaction-api_1.2_spec/1.0.1.Final/jboss-transaction-api_1.2_spec-1.0.1.Final.jar (28 kB at 142 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/common/hibernate-commons-annotations/5.0.1.Final/hibernate-commons-annotations-5.0.1.Final.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/dom4j/dom4j/1.6.1/dom4j-1.6.1.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jboss/jandex/2.0.3.Final/jandex-2.0.3.Final.jar (187 kB at 853 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/gitlab4j/gitlab4j-api/5.0.1/gitlab4j-api-5.0.1.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/fasterxml/classmate/1.3.0/classmate-1.3.0.jar (64 kB at 290 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/inject/jersey-hk2/2.35/jersey-hk2-2.35.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/common/hibernate-commons-annotations/5.0.1.Final/hibernate-commons-annotations-5.0.1.Final.jar (75 kB at 384 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-common/2.23.2/jersey-common-2.23.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/dom4j/dom4j/1.6.1/dom4j-1.6.1.jar (314 kB at 1.5 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-locator/2.5.0-b05/hk2-locator-2.5.0-b05.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/hibernate/hibernate-core/5.2.10.Final/hibernate-core-5.2.10.Final.jar (6.6 MB at 11 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/aopalliance-repackaged/2.5.0-b05/aopalliance-repackaged-2.5.0-b05.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/gitlab4j/gitlab4j-api/5.0.1/gitlab4j-api-5.0.1.jar (678 kB at 3.2 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-api/2.5.0-b05/hk2-api-2.5.0-b05.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/inject/jersey-hk2/2.35/jersey-hk2-2.35.jar (78 kB at 556 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-utils/2.5.0-b05/hk2-utils-2.5.0-b05.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/aopalliance-repackaged/2.5.0-b05/aopalliance-repackaged-2.5.0-b05.jar (15 kB at 111 kB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-locator/2.5.0-b05/hk2-locator-2.5.0-b05.jar (184 kB at 1.0 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-client/2.23.2/jersey-client-2.23.2.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/connectors/jersey-apache-connector/2.35/jersey-apache-connector-2.35.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-common/2.23.2/jersey-common-2.23.2.jar (715 kB at 3.6 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-multipart/2.35/jersey-media-multipart-2.35.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-utils/2.5.0-b05/hk2-utils-2.5.0-b05.jar (119 kB at 773 kB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/hk2-api/2.5.0-b05/hk2-api-2.5.0-b05.jar (179 kB at 940 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jvnet/mimepull/mimepull/1.9.13/mimepull-1.9.13.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-json-jackson/2.23.2/jersey-media-json-jackson-2.23.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/connectors/jersey-apache-connector/2.35/jersey-apache-connector-2.35.jar (46 kB at 374 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/ext/jersey-entity-filtering/2.23.2/jersey-entity-filtering-2.23.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-client/2.23.2/jersey-client-2.23.2.jar (169 kB at 1.2 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/jakarta/servlet/jakarta.servlet-api/4.0.4/jakarta.servlet-api-4.0.4.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-multipart/2.35/jersey-media-multipart-2.35.jar (82 kB at 450 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/uk/org/webcompere/system-stubs-jupiter/1.2.0/system-stubs-jupiter-1.2.0.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/jvnet/mimepull/mimepull/1.9.13/mimepull-1.9.13.jar (65 kB at 464 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/uk/org/webcompere/system-stubs-core/1.2.0/system-stubs-core-1.2.0.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-json-jackson/2.23.2/jersey-media-json-jackson-2.23.2.jar (22 kB at 111 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/chinamobile/cmos/sms-core/2.1.12.5/sms-core-2.1.12.5.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/jakarta/servlet/jakarta.servlet-api/4.0.4/jakarta.servlet-api-4.0.4.jar (83 kB at 638 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/de/ruedigermoeller/fst/2.48-jdk-6/fst-2.48-jdk-6.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/ext/jersey-entity-filtering/2.23.2/jersey-entity-filtering-2.23.2.jar (70 kB at 324 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/objenesis/objenesis/2.4/objenesis-2.4.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/uk/org/webcompere/system-stubs-jupiter/1.2.0/system-stubs-jupiter-1.2.0.jar (6.9 kB at 45 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/cedarsoftware/java-util/1.9.0/java-util-1.9.0.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/uk/org/webcompere/system-stubs-core/1.2.0/system-stubs-core-1.2.0.jar (46 kB at 257 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-transport/4.1.51.Final/netty-transport-4.1.51.Final.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/objenesis/objenesis/2.4/objenesis-2.4.jar (51 kB at 311 kB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/de/ruedigermoeller/fst/2.48-jdk-6/fst-2.48-jdk-6.jar (381 kB at 1.6 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-common/4.1.51.Final/netty-common-4.1.51.Final.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-buffer/4.1.51.Final/netty-buffer-4.1.51.Final.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/cedarsoftware/java-util/1.9.0/java-util-1.9.0.jar (58 kB at 342 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-resolver/4.1.51.Final/netty-resolver-4.1.51.Final.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-transport/4.1.51.Final/netty-transport-4.1.51.Final.jar (473 kB at 2.8 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec/4.1.51.Final/netty-codec-4.1.51.Final.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/chinamobile/cmos/sms-core/2.1.12.5/sms-core-2.1.12.5.jar (870 kB at 2.1 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-handler-proxy/4.1.51.Final/netty-handler-proxy-4.1.51.Final.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-common/4.1.51.Final/netty-common-4.1.51.Final.jar (625 kB at 4.0 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec-socks/4.1.51.Final/netty-codec-socks-4.1.51.Final.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-resolver/4.1.51.Final/netty-resolver-4.1.51.Final.jar (33 kB at 246 kB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-buffer/4.1.51.Final/netty-buffer-4.1.51.Final.jar (290 kB at 1.6 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec-http/4.1.51.Final/netty-codec-http-4.1.51.Final.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-handler/4.1.51.Final/netty-handler-4.1.51.Final.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec/4.1.51.Final/netty-codec-4.1.51.Final.jar (320 kB at 1.9 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/sleepycat/je/18.3.12/je-18.3.12.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-handler-proxy/4.1.51.Final/netty-handler-proxy-4.1.51.Final.jar (24 kB at 195 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/cglib/cglib-nodep/3.3.0/cglib-nodep-3.3.0.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec-socks/4.1.51.Final/netty-codec-socks-4.1.51.Final.jar (119 kB at 799 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-server/9.4.17.v20190418/jetty-server-9.4.17.v20190418.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-handler/4.1.51.Final/netty-handler-4.1.51.Final.jar (457 kB at 2.9 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-http/9.4.17.v20190418/jetty-http-9.4.17.v20190418.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/io/netty/netty-codec-http/4.1.51.Final/netty-codec-http-4.1.51.Final.jar (618 kB at 2.6 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-util/9.4.17.v20190418/jetty-util-9.4.17.v20190418.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/cglib/cglib-nodep/3.3.0/cglib-nodep-3.3.0.jar (415 kB at 1.7 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-io/9.4.17.v20190418/jetty-io-9.4.17.v20190418.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-http/9.4.17.v20190418/jetty-http-9.4.17.v20190418.jar (203 kB at 711 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-webapp/9.4.17.v20190418/jetty-webapp-9.4.17.v20190418.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/com/sleepycat/je/18.3.12/je-18.3.12.jar (3.5 MB at 7.9 MB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-server/9.4.17.v20190418/jetty-server-9.4.17.v20190418.jar (647 kB at 1.8 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-xml/9.4.17.v20190418/jetty-xml-9.4.17.v20190418.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-servlet/9.4.17.v20190418/jetty-servlet-9.4.17.v20190418.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-util/9.4.17.v20190418/jetty-util-9.4.17.v20190418.jar (526 kB at 1.7 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-security/9.4.17.v20190418/jetty-security-9.4.17.v20190418.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-io/9.4.17.v20190418/jetty-io-9.4.17.v20190418.jar (156 kB at 582 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/bundles/repackaged/jersey-guava/2.23.2/jersey-guava-2.23.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-webapp/9.4.17.v20190418/jetty-webapp-9.4.17.v20190418.jar (136 kB at 591 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/javax.inject/2.5.0-b05/javax.inject-2.5.0-b05.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-servlet/9.4.17.v20190418/jetty-servlet-9.4.17.v20190418.jar (121 kB at 572 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-servlet/2.23.2/jersey-container-servlet-2.23.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-xml/9.4.17.v20190418/jetty-xml-9.4.17.v20190418.jar (61 kB at 241 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-servlet-core/2.23.2/jersey-container-servlet-core-2.23.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/eclipse/jetty/jetty-security/9.4.17.v20190418/jetty-security-9.4.17.v20190418.jar (116 kB at 489 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-server/2.23.2/jersey-server-2.23.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/bundles/repackaged/jersey-guava/2.23.2/jersey-guava-2.23.2.jar (971 kB at 4.9 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-jaxb/2.23.2/jersey-media-jaxb-2.23.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/hk2/external/javax.inject/2.5.0-b05/javax.inject-2.5.0-b05.jar (6.0 kB at 38 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/providers/jersey-test-framework-provider-jdk-http/2.23.2/jersey-test-framework-provider-jdk-http-2.23.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-servlet/2.23.2/jersey-container-servlet-2.23.2.jar (18 kB at 82 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/jersey-test-framework-core/2.23.2/jersey-test-framework-core-2.23.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-servlet-core/2.23.2/jersey-container-servlet-core-2.23.2.jar (66 kB at 315 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/ow2/asm/asm-debug-all/5.0.4/asm-debug-all-5.0.4.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/providers/jersey-test-framework-provider-jdk-http/2.23.2/jersey-test-framework-provider-jdk-http-2.23.2.jar (7.1 kB at 38 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-jdk-http/2.23.2/jersey-container-jdk-http-2.23.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/media/jersey-media-jaxb/2.23.2/jersey-media-jaxb-2.23.2.jar (73 kB at 296 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-test/5.3.19/spring-test-5.3.19.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/core/jersey-server/2.23.2/jersey-server-2.23.2.jar (952 kB at 2.8 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-web/5.6.3/spring-security-web-5.6.3.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/test-framework/jersey-test-framework-core/2.23.2/jersey-test-framework-core-2.23.2.jar (30 kB at 104 kB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/ow2/asm/asm-debug-all/5.0.4/asm-debug-all-5.0.4.jar (379 kB at 1.4 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-core/5.6.3/spring-security-core-5.6.3.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-crypto/5.6.3/spring-security-crypto-5.6.3.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/spring-test/5.3.19/spring-test-5.3.19.jar (787 kB at 4.2 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/junit/vintage/junit-vintage-engine/5.8.2/junit-vintage-engine-5.8.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/glassfish/jersey/containers/jersey-container-jdk-http/2.23.2/jersey-container-jdk-http-2.23.2.jar (21 kB at 87 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/mockito/mockito-all/1.10.19/mockito-all-1.10.19.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-web/5.6.3/spring-security-web-5.6.3.jar (638 kB at 3.6 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-module-junit4/2.0.9/powermock-module-junit4-2.0.9.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/junit/vintage/junit-vintage-engine/5.8.2/junit-vintage-engine-5.8.2.jar (0 B at 0 B/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-module-junit4-common/2.0.9/powermock-module-junit4-common-2.0.9.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-crypto/5.6.3/spring-security-crypto-5.6.3.jar (82 kB at 474 kB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/springframework/security/spring-security-core/5.6.3/spring-security-core-5.6.3.jar (439 kB at 2.5 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-reflect/2.0.9/powermock-reflect-2.0.9.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/bytebuddy/byte-buddy/1.10.14/byte-buddy-1.10.14.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-module-junit4/2.0.9/powermock-module-junit4-2.0.9.jar (48 kB at 264 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/bytebuddy/byte-buddy-agent/1.10.14/byte-buddy-agent-1.10.14.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/mockito/mockito-all/1.10.19/mockito-all-1.10.19.jar (1.2 MB at 5.3 MB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-module-junit4-common/2.0.9/powermock-module-junit4-common-2.0.9.jar (18 kB at 98 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-core/2.0.9/powermock-core-2.0.9.jar
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/openjdk/jmh/jmh-core/1.35/jmh-core-1.35.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-reflect/2.0.9/powermock-reflect-2.0.9.jar (68 kB at 349 kB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/apache/commons/commons-math3/3.2/commons-math3-3.2.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/bytebuddy/byte-buddy-agent/1.10.14/byte-buddy-agent-1.10.14.jar (259 kB at 1.2 MB/s)
	//Downloading from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/openjdk/jmh/jmh-generator-annprocess/1.35/jmh-generator-annprocess-1.35.jar
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/openjdk/jmh/jmh-core/1.35/jmh-core-1.35.jar (541 kB at 2.4 MB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/net/bytebuddy/byte-buddy/1.10.14/byte-buddy-1.10.14.jar (3.5 MB at 10 MB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/powermock/powermock-core/2.0.9/powermock-core-2.0.9.jar (201 kB at 886 kB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/openjdk/jmh/jmh-generator-annprocess/1.35/jmh-generator-annprocess-1.35.jar (31 kB at 187 kB/s)
	//Downloaded from alimaven: http://maven.aliyun.com/nexus/content/groups/public/org/apache/commons/commons-math3/3.2/commons-math3-3.2.jar (1.7 MB at 6.1 MB/s)
	//[INFO] ------------------------------------------------------------------------
	//[INFO] BUILD SUCCESS
	//[INFO] ------------------------------------------------------------------------
	//[INFO] Total time:  37.065 s
	//[INFO] Finished at: 2023-06-27T22:28:35+08:00
	//[INFO] ------------------------------------------------------------------------

}
```

执行任意命令：

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
	// -c,--lax-checksums                     Warn if checksums don't match
	//    --color <arg>                       Defines the color mode of the
	//                                        output. Supported are 'auto',
	//                                        'always', 'never'.
	// -cpu,--check-plugin-updates            Ineffective, only kept for
	//                                        backward compatibility
	// -D,--define <arg>                      Define a system property
	// -e,--errors                            Produce execution error messages
	// -emp,--encrypt-master-password <arg>   Encrypt master security password
	// -ep,--encrypt-password <arg>           Encrypt server password
	// -f,--file <arg>                        Force the use of an alternate POM
	//                                        file (or directory with pom.xml)
	// -fae,--fail-at-end                     Only fail the build afterwards;
	//                                        allow all non-impacted builds to
	//                                        continue
	// -ff,--fail-fast                        Stop at first failure in
	//                                        reactorized builds
	// -fn,--fail-never                       NEVER fail the build, regardless
	//                                        of project result
	// -gs,--global-settings <arg>            Alternate path for the global
	//                                        settings file
	// -gt,--global-toolchains <arg>          Alternate path for the global
	//                                        toolchains file
	// -h,--help                              Display help information
	// -l,--log-file <arg>                    Log file where all build output
	//                                        will go (disables output color)
	// -llr,--legacy-local-repository         Use Maven 2 Legacy Local
	//                                        Repository behaviour, ie no use of
	//                                        _remote.repositories. Can also be
	//                                        activated by using
	//                                        -Dmaven.legacyLocalRepo=true
	// -N,--non-recursive                     Do not recurse into sub-projects
	// -npr,--no-plugin-registry              Ineffective, only kept for
	//                                        backward compatibility
	// -npu,--no-plugin-updates               Ineffective, only kept for
	//                                        backward compatibility
	// -nsu,--no-snapshot-updates             Suppress SNAPSHOT updates
	// -ntp,--no-transfer-progress            Do not display transfer progress
	//                                        when downloading or uploading
	// -o,--offline                           Work offline
	// -P,--activate-profiles <arg>           Comma-delimited list of profiles
	//                                        to activate
	// -pl,--projects <arg>                   Comma-delimited list of specified
	//                                        reactor projects to build instead
	//                                        of all projects. A project can be
	//                                        specified by [groupId]:artifactId
	//                                        or by its relative path
	// -q,--quiet                             Quiet output - only show errors
	// -rf,--resume-from <arg>                Resume reactor from specified
	//                                        project
	// -s,--settings <arg>                    Alternate path for the user
	//                                        settings file
	// -t,--toolchains <arg>                  Alternate path for the user
	//                                        toolchains file
	// -T,--threads <arg>                     Thread count, for instance 2.0C
	//                                        where C is core multiplied
	// -U,--update-snapshots                  Forces a check for missing
	//                                        releases and updated snapshots on
	//                                        remote repositories
	// -up,--update-plugins                   Ineffective, only kept for
	//                                        backward compatibility
	// -v,--version                           Display version information
	// -V,--show-version                      Display version information
	//                                        WITHOUT stopping build
	// -X,--debug                             Produce execution debug output

}
```



# 四、TODO

- 自动安装mvn
- API设计优化一下使其看起来像是给人用的...



