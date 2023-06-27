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
