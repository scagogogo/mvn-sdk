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
