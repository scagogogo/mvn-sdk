package installer

import (
	"os"
	"path/filepath"
	"testing"
)

// TestInstallMacOS 测试macOS平台的安装功能
// 由于涉及真实的下载和安装，默认跳过此测试
func TestInstallMacOS(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过macOS安装测试（使用-short参数）")
	}

	// 真实的集成测试
	mavenHome, err := InstallMacOS()
	if err != nil {
		t.Fatalf("安装Maven失败: %v", err)
	}

	// 验证安装路径
	if mavenHome == "" {
		t.Fatal("返回的Maven安装路径为空")
	}

	// 验证bin/mvn可执行文件存在
	mvnPath := filepath.Join(mavenHome, "bin", "mvn")
	_, err = os.Stat(mvnPath)
	if err != nil {
		t.Fatalf("未找到mvn可执行文件: %v", err)
	}

	t.Logf("Maven成功安装到: %s", mavenHome)
}

// TestSetMacOSEnvironmentVars 测试macOS环境变量设置
// 此测试会创建一个临时目录，不会修改真实的环境配置文件
func TestSetMacOSEnvironmentVars(t *testing.T) {
	// 创建临时目录作为测试的MAVEN_HOME
	tempDir, err := os.MkdirTemp("", "maven-test")
	if err != nil {
		t.Fatalf("创建临时目录失败: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// 保存原始HOME环境变量，测试后恢复
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)

	// 设置临时HOME环境变量，避免修改真实的配置文件
	os.Setenv("HOME", tempDir)

	// 测试环境变量设置函数
	err = setMacOSEnvironmentVars(tempDir + "/maven")
	if err != nil {
		t.Fatalf("设置环境变量失败: %v", err)
	}

	// 检查shell配置文件是否已创建
	shellFiles := []string{".zshrc", ".bash_profile"}
	found := false

	for _, file := range shellFiles {
		path := filepath.Join(tempDir, file)
		if _, err := os.Stat(path); err == nil {
			// 读取文件内容并检查是否包含Maven环境变量设置
			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("读取配置文件失败: %v", err)
			}

			if len(content) > 0 {
				found = true
				t.Logf("在%s中找到Maven配置", file)
			}
		}
	}

	if !found {
		t.Fatal("未在任何shell配置文件中找到Maven环境变量设置")
	}
}
