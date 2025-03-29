package installer

import (
	"fmt"
	"os"
	"path/filepath"
)

// 可测试版本的下载安装函数配置
type InstallOptions struct {
	// Maven下载URL
	MavenURL string
	// 用户目录（如果为空则使用真实的用户目录）
	HomeDir string
	// 是否跳过环境变量设置（测试时通常设为true）
	SkipEnvSetup bool
}

// 默认的安装选项
func DefaultInstallOptions() InstallOptions {
	return InstallOptions{
		MavenURL:     "https://dlcdn.apache.org/maven/maven-3/3.9.6/binaries/apache-maven-3.9.6-bin.tar.gz",
		HomeDir:      "",
		SkipEnvSetup: false,
	}
}

// 使用可配置选项安装Maven（macOS版本）
// 此函数供测试使用，允许注入依赖
func InstallMacOSWithOptions(options InstallOptions) (string, error) {
	// 确定用户目录
	homeDir := options.HomeDir
	if homeDir == "" {
		var err error
		homeDir, err = os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("获取用户主目录失败: %w", err)
		}
	}

	// 创建安装目录
	mavenDir := filepath.Join(homeDir, ".m2", "maven")
	if err := os.MkdirAll(mavenDir, 0755); err != nil {
		return "", fmt.Errorf("创建Maven安装目录失败: %w", err)
	}

	// 下载Maven
	tarPath := filepath.Join(mavenDir, "maven.tar.gz")
	if err := downloadFile(options.MavenURL, tarPath); err != nil {
		return "", fmt.Errorf("下载Maven失败: %w", err)
	}

	// 解压Maven
	extractDir := filepath.Join(mavenDir, "maven-install")
	if err := untar(tarPath, extractDir); err != nil {
		return "", fmt.Errorf("解压Maven失败: %w", err)
	}

	// 查找解压后的目录
	mavenHome := ""
	err := filepath.Walk(extractDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && filepath.Base(path) == "apache-maven-3.9.6" {
			mavenHome = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("查找Maven目录失败: %w", err)
	}

	if mavenHome == "" {
		return "", fmt.Errorf("找不到Maven安装目录")
	}

	// 确保bin目录存在
	binDir := filepath.Join(mavenHome, "bin")
	if _, err := os.Stat(binDir); os.IsNotExist(err) {
		return "", fmt.Errorf("Maven安装不完整，找不到bin目录")
	}

	// 设置环境变量（可跳过）
	if !options.SkipEnvSetup {
		if err := setMacOSEnvironmentVars(mavenHome); err != nil {
			return "", fmt.Errorf("设置环境变量失败: %w", err)
		}
	}

	// 清理临时文件
	os.Remove(tarPath)

	return mavenHome, nil
}
