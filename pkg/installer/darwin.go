package installer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// InstallMacOS 在macOS上安装Maven
func InstallMacOS() (string, error) {
	// 尝试使用Homebrew安装
	if installed, path := tryHomebrewInstall(); installed {
		return path, nil
	}

	// 如果Homebrew安装失败，则使用二进制包安装（类似Linux）
	return installFromBinaryMacOS()
}

// 通过Homebrew安装Maven
func tryHomebrewInstall() (bool, string) {
	// 检查是否有brew命令
	cmd := exec.Command("which", "brew")
	if err := cmd.Run(); err == nil {
		// 使用brew安装maven
		cmd = exec.Command("brew", "install", "maven")
		if err := cmd.Run(); err == nil {
			// 安装成功，查找安装路径
			cmd = exec.Command("which", "mvn")
			output, err := cmd.Output()
			if err == nil && len(output) > 0 {
				mvnPath := strings.TrimSpace(string(output))
				// 获取MAVEN_HOME
				mavenHome := filepath.Dir(filepath.Dir(mvnPath))
				return true, mavenHome
			}
		}
	}
	return false, ""
}

// 从二进制包安装macOS版本的Maven
func installFromBinaryMacOS() (string, error) {
	// 使用与Linux相同的tar.gz包
	mavenURL := "https://dlcdn.apache.org/maven/maven-3/3.9.6/binaries/apache-maven-3.9.6-bin.tar.gz"

	mavenHome, err := installFromTarGz(mavenURL)
	if err != nil {
		return "", err
	}

	// 为macOS设置环境变量
	if err := setMacOSEnvironmentVars(mavenHome); err != nil {
		return "", fmt.Errorf("设置环境变量失败: %w", err)
	}

	return mavenHome, nil
}

// 设置macOS环境变量
func setMacOSEnvironmentVars(mavenHome string) error {
	// 向.zshrc或.bash_profile写入环境变量
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// 检查用户使用的shell
	var rcFile string
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "zsh") {
		rcFile = filepath.Join(homeDir, ".zshrc")
	} else {
		// 默认使用bash
		rcFile = filepath.Join(homeDir, ".bash_profile")
	}

	// 读取现有文件内容
	content, err := os.ReadFile(rcFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	// 构建新的环境变量设置
	envVars := fmt.Sprintf("\n# Maven环境变量\nexport MAVEN_HOME=%s\nexport PATH=$PATH:$MAVEN_HOME/bin\n", mavenHome)

	// 检查是否已经有这些设置
	if !strings.Contains(string(content), "MAVEN_HOME=") {
		// 追加到文件
		f, err := os.OpenFile(rcFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.WriteString(envVars); err != nil {
			return err
		}
	}

	fmt.Println("Maven环境变量已设置，请运行 'source " + rcFile + "' 或重新打开终端使其生效")

	return nil
}
