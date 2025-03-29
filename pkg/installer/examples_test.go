package installer

import (
	"fmt"
)

// ExampleInstall 展示如何使用Install函数安装Maven
func Example_install() {
	// 只是展示用法，不会实际运行
	fmt.Println("Maven已成功安装到: /path/to/maven")
	fmt.Println("可以使用以下命令验证安装:")
	fmt.Println("/path/to/maven/bin/mvn --version")

	// Output:
	// Maven已成功安装到: /path/to/maven
	// 可以使用以下命令验证安装:
	// /path/to/maven/bin/mvn --version
}

// 测试说明
func Example_testInstructions() {
	fmt.Println("如何运行这些测试:")
	fmt.Println("1. 运行所有测试，包括集成测试（可能会实际安装Maven）:")
	fmt.Println("   go test -v ./pkg/installer")
	fmt.Println()
	fmt.Println("2. 仅运行单元测试，跳过集成测试:")
	fmt.Println("   go test -v -short ./pkg/installer")
	fmt.Println()
	fmt.Println("3. 运行特定测试（例如，只运行下载测试）:")
	fmt.Println("   go test -v -run TestDownloadFile ./pkg/installer")
	fmt.Println()
	fmt.Println("4. 运行模拟安装测试:")
	fmt.Println("   go test -v -run TestInstallMacOSWithMock ./pkg/installer")

	// Output:
	// 如何运行这些测试:
	// 1. 运行所有测试，包括集成测试（可能会实际安装Maven）:
	//    go test -v ./pkg/installer
	//
	// 2. 仅运行单元测试，跳过集成测试:
	//    go test -v -short ./pkg/installer
	//
	// 3. 运行特定测试（例如，只运行下载测试）:
	//    go test -v -run TestDownloadFile ./pkg/installer
	//
	// 4. 运行模拟安装测试:
	//    go test -v -run TestInstallMacOSWithMock ./pkg/installer
}

// 本地安装Maven的示例
func Example_macOSInstall() {
	// 只是展示用法，不会实际运行
	fmt.Println("在macOS上安装Maven:")
	fmt.Println("Maven已成功安装到: /Users/username/.m2/maven/maven-install/apache-maven-3.9.6")
	fmt.Println("您可以通过在终端中运行以下命令验证安装:")
	fmt.Println("mvn -v")

	// Output:
	// 在macOS上安装Maven:
	// Maven已成功安装到: /Users/username/.m2/maven/maven-install/apache-maven-3.9.6
	// 您可以通过在终端中运行以下命令验证安装:
	// mvn -v
}
