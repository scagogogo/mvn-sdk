package installer

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// InstallLinux 在Linux系统上安装Maven
func InstallLinux() (string, error) {
	// 尝试使用包管理器安装
	if installed, path := tryPackageManagerLinux(); installed {
		return path, nil
	}

	// 如果包管理器安装失败，则使用二进制包安装
	return installFromBinaryLinux()
}

// 尝试使用包管理器安装
func tryPackageManagerLinux() (bool, string) {
	// 检测系统类型
	if _, err := os.Stat("/etc/debian_version"); err == nil {
		// Debian/Ubuntu系统
		cmd := exec.Command("sudo", "apt-get", "update")
		cmd.Run()

		cmd = exec.Command("sudo", "apt-get", "install", "-y", "maven")
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
	} else if _, err := os.Stat("/etc/redhat-release"); err == nil {
		// RedHat/CentOS/Fedora系统
		cmd := exec.Command("sudo", "yum", "install", "-y", "maven")
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

// 从二进制包安装
func installFromBinaryLinux() (string, error) {
	// Maven下载地址
	mavenURL := "https://dlcdn.apache.org/maven/maven-3/3.9.6/binaries/apache-maven-3.9.6-bin.tar.gz"
	// 安装目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("获取用户主目录失败: %w", err)
	}

	// 创建安装目录
	mavenDir := filepath.Join(homeDir, ".m2", "maven")
	if err := os.MkdirAll(mavenDir, 0755); err != nil {
		return "", fmt.Errorf("创建Maven安装目录失败: %w", err)
	}

	// 下载Maven
	tarPath := filepath.Join(mavenDir, "maven.tar.gz")
	if err := downloadFileLinux(mavenURL, tarPath); err != nil {
		return "", fmt.Errorf("下载Maven失败: %w", err)
	}

	// 解压Maven
	extractDir := filepath.Join(mavenDir, "maven-install")
	if err := untarLinux(tarPath, extractDir); err != nil {
		return "", fmt.Errorf("解压Maven失败: %w", err)
	}

	// 查找解压后的目录
	mavenHome := ""
	err = filepath.Walk(extractDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && strings.Contains(info.Name(), "apache-maven") {
			mavenHome = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("查找Maven目录失败: %w", err)
	}

	if mavenHome == "" {
		return "", errors.New("找不到Maven安装目录")
	}

	// 确保bin目录存在
	binDir := filepath.Join(mavenHome, "bin")
	if _, err := os.Stat(binDir); os.IsNotExist(err) {
		return "", errors.New("Maven安装不完整，找不到bin目录")
	}

	// 设置环境变量
	if err := setEnvironmentVarsLinux(mavenHome); err != nil {
		return "", fmt.Errorf("设置环境变量失败: %w", err)
	}

	// 删除临时文件
	os.Remove(tarPath)

	return mavenHome, nil
}

// 下载文件
func downloadFileLinux(url, destPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，HTTP状态码: %d", resp.StatusCode)
	}

	file, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

// 解压tar.gz文件
func untarLinux(tarPath, destDir string) error {
	file, err := os.Open(tarPath)
	if err != nil {
		return err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	// 创建目标目录
	if err := os.MkdirAll(destDir, 0755); err != nil {
		return err
	}

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		path := filepath.Join(destDir, header.Name)

		// 检查路径穿越漏洞
		if !strings.HasPrefix(path, filepath.Clean(destDir)+string(os.PathSeparator)) {
			return fmt.Errorf("非法的文件路径: %s", path)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			dir := filepath.Dir(path)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}

			outFile, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}

	return nil
}

// 设置环境变量
func setEnvironmentVarsLinux(mavenHome string) error {
	// 向.bashrc或.zshrc写入环境变量
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
		rcFile = filepath.Join(homeDir, ".bashrc")
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
